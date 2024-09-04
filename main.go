package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var subvolumeNames []string

var pvFromKubeArr []string

var LOG_LEVEL = getEnv("LOG_LEVEL", "INFO")
var KUBE_CONFIG_FILE_PATH = os.Getenv("KUBE_CONFIG_FILE_PATH")

// LogEntry logs in json mod
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

// getEnv take default value if env not exist
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logWithTime(fmt.Sprintf("Couldn't find external var %s. Use %s by default.🦄", key, defaultValue))
	return defaultValue
}

func logWithTime(message string) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   message,
	}
	jsonData, err := json.Marshal(entry)
	if err != nil {
		log.Fatalf("Failed to marshal log entry to JSON: %v", err)
		return
	}

	fmt.Println(string(jsonData))
}

func main() {
	// Загружаем список из файла
	// Путь к файлу
	filePath := "subvolumeListFronCeph.txt"

	// Загружаем имена из файла
	err := loadSubvolumeNamesFromFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
		return
	}

	// Выводим загруженные значения
	if LOG_LEVEL == "DEBUG" {
		logWithTime(fmt.Sprintf("Array from file: %+v", subvolumeNames))
	}

	var kubeconfig *string

	if KUBE_CONFIG_FILE_PATH == "" {
		// Переменная окружения не задана
		fmt.Println("Переменная окружения не задана. Выполняем действие по умолчанию.")
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
	} else {
		// Переменная окружения задана
		fmt.Printf("Переменная окружения задана: %s\n", KUBE_CONFIG_FILE_PATH)

		// Определяем флаг для kubeconfig

		defaultKubeConfigPath := KUBE_CONFIG_FILE_PATH

		// Указываем флаг для возможности переопределить путь к файлу конфигурации
		kubeconfig = flag.String("kubeconfig", defaultKubeConfigPath, "(опционально) абсолютный путь до kubeconfig файла")

		flag.Parse()

	}

	// Создание конфигурации клиента
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Создание клиента для Kubernetes
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Получение списка PV
	pvList, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing PVs: %v", err)
	}

	// Печать информации о PV
	for _, pv := range pvList.Items {
		findPv(pv)
	}

	findMatchesAndDifferences(subvolumeNames, pvFromKubeArr)

}

func findPv(pv v1.PersistentVolume) {
	capacity := pv.Spec.Capacity[v1.ResourceStorage]
	capacityStr := capacity.String()

	// Проверка, что PV использует CSI драйвер
	if pv.Spec.CSI != nil {
		subvolumeName := pv.Spec.CSI.VolumeAttributes["subvolumeName"]

		if LOG_LEVEL == "DEBUG" {
			message := fmt.Sprintf("Find PV in k8s. Name: %s Capacity: %s Access Modes: %v SubvolumeName: %s", pv.Name, capacityStr, pv.Spec.AccessModes, subvolumeName)
			logWithTime(message)
		}

		pvFromKubeArr = append(pvFromKubeArr, subvolumeName)

	} else {

		message := fmt.Sprintf("Get PV from k8s. Name: %sCapacity: %sAccess Modes: %vNo CSI driver info available", pv.Name, capacityStr, pv.Spec.AccessModes)
		logWithTime(message)
	}

	if LOG_LEVEL == "DEBUG" {
		logWithTime(fmt.Sprintf("PersistentVolumes array: %+v", pvFromKubeArr))
	}

}

func findMatchesAndDifferences(list1, list2 []string) {
	matches := []string{}
	differences := []string{}
	seen := make(map[string]int)

	// Заполняем карту, отмечая элементы из обоих списков
	for _, item := range list1 {
		seen[item]++
	}

	for _, item := range list2 {
		seen[item]--
	}

	// Ищем совпадения и различия
	for item, count := range seen {
		if count == 0 {
			matches = append(matches, item)
		} else {
			differences = append(differences, item)
		}
	}

	// Выводим результаты
	logWithTime(fmt.Sprintf("🌽Matches %v values: %+v", len(matches), matches))
	// logWithTime(fmt.Sprintf("🦐Differences %v values: %+v", len(differences), differences))
}

func loadSubvolumeNamesFromFile(filePath string) error {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Используем буферизованный сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			// Удаляем запятые и кавычки
			line = strings.Trim(line, `",`)
			subvolumeNames = append(subvolumeNames, line)
		}
	}

	// Проверка на ошибки чтения файла
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	return nil
}
