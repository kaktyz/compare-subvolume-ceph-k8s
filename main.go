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

//	var subvolumeNames = []string{
//		"csi-vol-e61ab130-4fe6-11ef-a388-c6f2daeeb83a",
//		"csi-vol-95a72863-4438-11ef-a388-c6f2daeeb83a",
//		"csi-vol-cc82874e-2fc3-11ed-b12f-8af21f51cb6d",
//		"csi-vol-d6b0ff74-f4b4-11ed-8198-7a0d435175b5",
//		"csi-vol-f034e76f-0f5f-11ee-a9fd-3ebe5e805178",
//		"csi-vol-a5d1a70b-efdc-11ed-b280-9e559261c936",
//		"csi-vol-cfe493c0-2d46-11ef-9f23-fe3532ece709",
//		"csi-vol-9afc87d0-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-fc9749a1-4957-11ed-ba52-c69a66c45312",
//		"csi-vol-2e933111-756a-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-6eaaf12f-391c-4f4c-b2f6-d10946e56330",
//		"csi-vol-24058251-89eb-11ee-9abc-ea01a90ca1db",
//		"csi-vol-95c5d111-4438-11ef-a388-c6f2daeeb83a",
//		"csi-vol-bc818f7f-69ab-11ed-8233-9e469fa71e3b",
//		"csi-vol-f0a06ccf-6f21-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-6686b51d-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-3299a49f-adf5-11ed-b9fc-7a55e2a23421",
//		"csi-vol-d65e5346-0527-11ee-a9fd-3ebe5e805178",
//		"csi-vol-894da1f1-95a3-11ee-a30a-92d7948fed0a",
//		"csi-vol-75c842f0-ca2b-11ed-9bbe-baccb3867649",
//		"csi-vol-6449f16e-2fd1-11ed-b12f-8af21f51cb6d",
//		"csi-vol-5e5a9bee-2cfd-11ed-9914-428aff3909c1",
//		"csi-vol-2a3a8dbb-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-9afa7f5d-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-95fa1292-e027-11ec-a257-764b17edad84",
//		"csi-vol-4d9d810d-efd4-11ed-b280-9e559261c936",
//		"csi-vol-19f60227-f3c1-11ed-b7fb-da6452cd87b3",
//		"csi-vol-3c94fd82-0080-11ef-9675-c2dd098820e2",
//		"csi-vol-bf4857b6-66a1-11ee-9abc-ea01a90ca1db",
//		"csi-vol-27b98e13-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-43378e61-03c6-11ee-9b40-caf255bd9603",
//		"csi-vol-e1c9d71c-49be-11ef-a388-c6f2daeeb83a",
//		"csi-vol-e61a1433-4fe6-11ef-a388-c6f2daeeb83a",
//		"csi-vol-cc248b7d-039b-11ee-be49-be8a2a532096",
//		"csi-vol-1e48b5dd-f090-11ed-b280-9e559261c936",
//		"csi-vol-5a795f5d-2fc6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-0dfddaab-199b-11ed-a13e-2293c23da942",
//		"csi-vol-bfcc82ea-607a-11ef-9e71-66877a12684d",
//		"csi-vol-53bd3fc5-95a3-11ee-a30a-92d7948fed0a",
//		"csi-vol-25efbc89-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-f4078fcc-bb69-11ee-a30a-92d7948fed0a",
//		"csi-vol-449f59ca-2fce-11ed-b12f-8af21f51cb6d",
//		"csi-vol-555c3fbc-2fb6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-7cc33721-efdb-11ed-b280-9e559261c936",
//		"csi-vol-ec0023ec-69a8-11ed-8233-9e469fa71e3b",
//		"csi-vol-28addf57-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-21d71b44-2fb6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-76d6c831-efdc-11ed-b280-9e559261c936",
//		"csi-vol-eea5f5c4-2fc2-11ed-b12f-8af21f51cb6d",
//		"csi-vol-398d6050-4066-11ef-a388-c6f2daeeb83a",
//		"csi-vol-5a66a62a-2cfd-11ed-9914-428aff3909c1",
//		"csi-vol-27210a50-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-a066c0dc-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-b1532f3e-3484-11ef-a388-c6f2daeeb83a",
//		"csi-vol-f549e60f-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-5e34486c-7faf-11ed-85b1-564f006a9578",
//		"csi-vol-5be09b68-22f7-48c1-be89-2a80d4e1b632",
//		"csi-vol-89f083ec-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-c819952e-493e-11ed-ba52-c69a66c45312",
//		"csi-vol-8659b54f-e769-4ac6-816a-9d69f609b333",
//		"csi-vol-16316413-443d-11ef-a388-c6f2daeeb83a",
//		"csi-vol-9afb08ba-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-71fba888-2560-11ee-ae23-6e9adfdafb04",
//		"csi-vol-ee457cca-efdc-11ed-b280-9e559261c936",
//		"csi-vol-2b6ba3d4-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-6fe3dcdc-39e0-11ef-bd0a-2ec837d2743f",
//		"csi-vol-a4dcaf4c-d394-4da7-92c4-99ffbd3219cd",
//		"csi-vol-163177a9-443d-11ef-a388-c6f2daeeb83a",
//		"csi-vol-1915a27d-2fb4-11ed-b12f-8af21f51cb6d",
//		"csi-vol-ae72a159-6925-11ef-9f88-e60b227d963c",
//		"csi-vol-ee48387c-efdc-11ed-b280-9e559261c936",
//		"csi-vol-5e23f91a-7faf-11ed-85b1-564f006a9578",
//		"csi-vol-50957286-efd5-11ed-b280-9e559261c936",
//		"csi-vol-2ab4a50f-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-e8d1ccbc-1a41-11ed-b73f-7adb244cf5bc",
//		"csi-vol-9b73e0a9-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-0e088292-199b-11ed-a13e-2293c23da942",
//		// "csi-vol-f0c21a98-6929-11ef-9f88-e60b227d963c",
//		"csi-vol-72a5f46a-4070-11ef-a388-c6f2daeeb83a",
//		"csi-vol-42ed5113-e02b-11ec-a257-764b17edad84",
//		"csi-vol-2eb21d53-1a2e-11ed-b73f-7adb244cf5bc",
//		"csi-vol-097327ce-2548-11ee-84d6-3eee9826907e",
//		"csi-vol-19638792-57e1-4678-8714-97c1a33a1cfe",
//		"csi-vol-db83aa30-d9f6-11ed-aaf6-867bdcd35510",
//		"csi-vol-f1ff743b-6929-11ef-9f88-e60b227d963c",
//		"csi-vol-fc7e8f3d-4957-11ed-ba52-c69a66c45312",
//		"csi-vol-8c5d04a5-443e-11ef-a388-c6f2daeeb83a",
//		"csi-vol-6037e092-54a8-11ef-9e71-66877a12684d",
//		"csi-vol-9afa8546-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-567a86db-4a10-11ed-ba52-c69a66c45312",
//		"csi-vol-494b53a0-2cf6-11ed-9914-428aff3909c1",
//		"csi-vol-0004badd-efd4-11ed-b280-9e559261c936",
//		"csi-vol-775af577-adf5-11ed-b9fc-7a55e2a23421",
//		"csi-vol-3a02d031-2099-11ee-95eb-cad15fd723cb",
//		"csi-vol-77c3de28-4554-4135-991b-6373b27242ef",
//		"csi-vol-24138687-89eb-11ee-9abc-ea01a90ca1db",
//		"csi-vol-255485d7-4064-11ef-a388-c6f2daeeb83a",
//		"csi-vol-86855081-f2fe-11ed-b7fb-da6452cd87b3",
//		"csi-vol-507788c0-1b67-11ee-a9fd-3ebe5e805178",
//		"csi-vol-c9ba7ec8-2196-11ef-8c56-d642877ef82c",
//		"csi-vol-a0ef7d13-77d1-11ee-9abc-ea01a90ca1db",
//		"csi-vol-7e3095b8-e998-11ed-bc7f-ea4b64fd219f",
//		"csi-vol-2e9ef60f-756a-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-2ad2ff83-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-fb4df034-fd66-4549-bf04-5504dbb6e89a",
//		"csi-vol-f03aa702-0f5f-11ee-a9fd-3ebe5e805178",
//		"csi-vol-6cdc9f39-5559-11ef-9e43-ca882da5a549",
//		"csi-vol-8004ed39-adf5-11ed-b9fc-7a55e2a23421",
//		"csi-vol-29fd63ce-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-1bc4791e-2fc3-11ed-b12f-8af21f51cb6d",
//		"csi-vol-9c114c54-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-9bb135ed-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-65c2d3a3-e501-11ed-b7fb-da6452cd87b3",
//		"csi-vol-1ce68387-c7f1-11ed-9bbe-baccb3867649",
//		"csi-vol-2525ab32-42a2-11ef-a388-c6f2daeeb83a",
//		"csi-vol-2964c069-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-5bfdcfaf-7faf-11ed-85b1-564f006a9578",
//		"csi-vol-27d7e96d-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-548821be-7aa5-47e2-ab46-096d1189a820",
//		"csi-vol-bf4bae7e-42b2-11ef-a388-c6f2daeeb83a",
//		"csi-vol-16ef7842-2f9c-11ed-b12f-8af21f51cb6d",
//		"csi-vol-2a95e62c-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-2e89b41d-756a-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-9b92ba2b-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-1e4aabf7-0aa3-11ee-9119-6691476ab1bd",
//		"csi-vol-868ae061-2fce-11ed-b12f-8af21f51cb6d",
//		"csi-vol-a5649acd-efdc-11ed-b280-9e559261c936",
//		"csi-vol-250b5a67-42a2-11ef-a388-c6f2daeeb83a",
//		"csi-vol-c3b1ab52-2fcd-11ed-b12f-8af21f51cb6d",
//		"csi-vol-2eb21fae-1a2e-11ed-b73f-7adb244cf5bc",
//		"csi-vol-6104958a-fec1-11ed-9b40-caf255bd9603",
//		"csi-vol-ce7827e8-2fb5-11ed-b12f-8af21f51cb6d",
//		"csi-vol-a1daad35-b1bc-11ed-9bbe-baccb3867649",
//		"csi-vol-6ea1ea49-efd7-11ed-b280-9e559261c936",
//		"csi-vol-86dba8b6-eddd-465b-9bc1-16ca07d380aa",
//		"csi-vol-9bcfd5b8-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-55502910-2fb6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-d66a4dd7-0527-11ee-a9fd-3ebe5e805178",
//		"csi-vol-e67d5575-4433-11ef-a388-c6f2daeeb83a",
//		"csi-vol-37c6ba4a-69c3-11ef-9e43-ca882da5a549",
//		"csi-vol-2b4d5e3c-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-60eb0c61-fec1-11ed-9b40-caf255bd9603",
//		"csi-vol-f04c0aed-6f21-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-d192d269-2fc8-11ed-b12f-8af21f51cb6d",
//		"csi-vol-190021c4-2fb4-11ed-b12f-8af21f51cb6d",
//		"csi-vol-fc68afc1-efd0-11ed-b280-9e559261c936",
//		"csi-vol-53cda5ed-1f37-11ef-a388-c6f2daeeb83a",
//		"csi-vol-45c5403b-2fc6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-f1c5ef29-4437-11ef-a388-c6f2daeeb83a",
//		"csi-vol-867324fb-f2fe-11ed-b7fb-da6452cd87b3",
//		"csi-vol-288f335c-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-bf4ba8be-42b2-11ef-a388-c6f2daeeb83a",
//		"csi-vol-ef6e1c41-2fce-11ed-b12f-8af21f51cb6d",
//		"csi-vol-7b885297-871c-4c83-96bc-1aa11a1cda91",
//		"csi-vol-5ef44b78-53f2-11ef-9f88-e60b227d963c",
//		"csi-vol-024938aa-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-0474c140-4884-11ed-ba52-c69a66c45312",
//		"csi-vol-023f32dd-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-9afac668-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-a56b090e-efdc-11ed-b280-9e559261c936",
//		"csi-vol-9afeef4e-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-cbf06d9a-f088-4d52-92e1-552c6dcfe1c0",
//		"csi-vol-42eeed6e-e02b-11ec-a257-764b17edad84",
//		"csi-vol-956df0cc-66c9-11ef-9f88-e60b227d963c",
//		"csi-vol-b8813856-7aeb-11ed-985b-7a4a83bff840",
//		"csi-vol-b217aeab-2fb4-11ed-b12f-8af21f51cb6d",
//		"csi-vol-9afeff75-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-64f84528-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-1bccb4f0-2fc3-11ed-b12f-8af21f51cb6d",
//		"csi-vol-0e02b07d-199b-11ed-a13e-2293c23da942",
//		"csi-vol-57c60f77-42a2-11ef-a388-c6f2daeeb83a",
//		"csi-vol-cc728569-5896-11ef-9e43-ca882da5a549",
//		"csi-vol-29839198-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-989131c3-0603-11ee-a9fd-3ebe5e805178",
//		"csi-vol-e9cdfcac-441f-11ef-a388-c6f2daeeb83a",
//		"csi-vol-eebb3d31-2fc2-11ed-b12f-8af21f51cb6d",
//		"csi-vol-f2a07ecd-efda-11ed-b280-9e559261c936",
//		"csi-vol-1f31c3fb-2eb4-4780-ad5e-b240c2908e74",
//		"csi-vol-2facb7b1-d0a4-4b4a-b723-ff5764b91cd4",
//		"csi-vol-47455f30-efd2-11ed-b280-9e559261c936",
//		"csi-vol-5bd525fb-7faf-11ed-85b1-564f006a9578",
//		"csi-vol-f046a5da-7d6f-11ee-9abc-ea01a90ca1db",
//		"csi-vol-573b5dfe-f938-11ed-9b40-caf255bd9603",
//		"csi-vol-50c4a2ab-d2a5-47d9-b27b-63074ece9789",
//		"csi-vol-27029e12-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-f0417d06-0f5f-11ee-a9fd-3ebe5e805178",
//		"csi-vol-64e546ca-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-30c8a4b8-69e3-11ef-9e43-ca882da5a549",
//		"csi-vol-e67d8b3c-4433-11ef-a388-c6f2daeeb83a",
//		"csi-vol-0dfb43ec-199b-11ed-a13e-2293c23da942",
//		"csi-vol-1cf2372d-efdc-11ed-b280-9e559261c936",
//		"csi-vol-4393cee0-2fc7-11ed-b12f-8af21f51cb6d",
//		"csi-vol-a899e0f5-069f-11ee-9119-6691476ab1bd",
//		"csi-vol-3e8f71b5-336f-11ed-b12f-8af21f51cb6d",
//		"csi-vol-f1e3e819-4437-11ef-a388-c6f2daeeb83a",
//		"csi-vol-8771cdba-2fb4-11ed-b12f-8af21f51cb6d",
//		"csi-vol-ce7102f6-2fb5-11ed-b12f-8af21f51cb6d",
//		"csi-vol-bfcc8ee1-607a-11ef-9e71-66877a12684d",
//		"csi-vol-6f42f6f7-4da4-11ef-a388-c6f2daeeb83a",
//		"csi-vol-12c17d31-569e-11ed-b12f-8af21f51cb6d",
//		"csi-vol-ec07d867-69a8-11ed-8233-9e469fa71e3b",
//		"csi-vol-b5a5c6d2-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-b7ea87b6-5fb0-11ef-9e71-66877a12684d",
//		"csi-vol-9bee4fc2-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-53d24434-95a3-11ee-a30a-92d7948fed0a",
//		"csi-vol-7913db6b-36b9-11ee-9b52-82b82dddcb3e",
//		"csi-vol-1a0c44e1-f3c1-11ed-b7fb-da6452cd87b3",
//		"csi-vol-1d197417-597d-49db-af70-741d9fbf0fbd",
//		"csi-vol-2eb220f0-1a2e-11ed-b73f-7adb244cf5bc",
//		"csi-vol-f1315610-6f21-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-568db798-4a10-11ed-ba52-c69a66c45312",
//		"csi-vol-fa811ffa-1011-4674-b579-a260dfd1098b",
//		"csi-vol-678a9907-0080-11ef-9675-c2dd098820e2",
//		"csi-vol-4bfb8248-9859-46b1-944a-29aa0cbc1e2a",
//		"csi-vol-eb40831e-42af-11ef-a388-c6f2daeeb83a",
//		"csi-vol-f0f412b1-6f21-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-ca6d916d-efdc-11ed-b280-9e559261c936",
//		"csi-vol-e7f872a5-efd7-11ed-b280-9e559261c936",
//		"csi-vol-db8e45ca-d9f6-11ed-aaf6-867bdcd35510",
//		"csi-vol-25b2ea04-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-cc6e3cc0-2fc3-11ed-b12f-8af21f51cb6d",
//		"csi-vol-18612028-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-393cf0c6-20e4-4f74-9009-b1098d4888a4",
//		"csi-vol-c8024d18-493e-11ed-ba52-c69a66c45312",
//		"csi-vol-9c2b96e5-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-4db2c02d-efd4-11ed-b280-9e559261c936",
//		"csi-vol-3b0c400c-2099-11ee-95eb-cad15fd723cb",
//		"csi-vol-fc587729-efd0-11ed-b280-9e559261c936",
//		"csi-vol-494831f0-2cf6-11ed-9914-428aff3909c1",
//		"csi-vol-cc5bf5f1-1f46-11ef-a388-c6f2daeeb83a",
//		"csi-vol-a1dd6a25-b1bc-11ed-9bbe-baccb3867649",
//		"csi-vol-c52e19df-2d4f-11ef-bee1-061aeeda9862",
//		"csi-vol-2eb290f3-1a2e-11ed-b73f-7adb244cf5bc",
//		"csi-vol-8961e5b8-95a3-11ee-a30a-92d7948fed0a",
//		"csi-vol-57b25f6f-42a2-11ef-a388-c6f2daeeb83a",
//		"csi-vol-f202e3fc-4437-11ef-a388-c6f2daeeb83a",
//		"csi-vol-001e56f2-efd4-11ed-b280-9e559261c936",
//		"csi-vol-f53300d2-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-0df92985-199b-11ed-a13e-2293c23da942",
//		"csi-vol-532a8f42-340c-11ed-b12f-8af21f51cb6d",
//		"csi-vol-530eb501-95a3-11ee-a30a-92d7948fed0a",
//		"csi-vol-df5159b5-3c28-4312-9d17-5f2f17c499c1",
//		"csi-vol-e8d9d8e1-1a41-11ed-b73f-7adb244cf5bc",
//		"csi-vol-a8ebcdd9-11f7-49b0-b030-78f288f6b716",
//		"csi-vol-a0c93c1a-5930-11ee-9b52-82b82dddcb3e",
//		"csi-vol-7763e70d-adf5-11ed-b9fc-7a55e2a23421",
//		"csi-vol-29a1e532-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-7eecc009-6429-11ee-9abc-ea01a90ca1db",
//		"csi-vol-af98a6e8-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-0d4bfbe9-9d87-11ee-a30a-92d7948fed0a",
//		"csi-vol-7575b33d-beed-11ec-8db4-5abd84c443b3",
//		"csi-vol-0403921c-2d44-11ef-9f23-fe3532ece709",
//		"csi-vol-87943cb6-dfc9-4b4e-84c2-325ecb3835dd",
//		"csi-vol-645fdba0-2fd1-11ed-b12f-8af21f51cb6d",
//		"csi-vol-c55a802a-2fcd-11ed-b12f-8af21f51cb6d",
//		"csi-vol-c15efb39-4067-11ef-a388-c6f2daeeb83a",
//		"csi-vol-2b2e8a0d-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-bc6bec65-69ab-11ed-8233-9e469fa71e3b",
//		"csi-vol-a7db884f-1189-4259-b11c-3b32f6983b90",
//		"csi-vol-c15efd0b-4067-11ef-a388-c6f2daeeb83a",
//		"csi-vol-150b0c7e-1f37-11ef-a388-c6f2daeeb83a",
//		"csi-vol-2b47038f-07ad-495f-8710-259907de231f",
//		"csi-vol-e7be72cc-ef13-11ed-b280-9e559261c936",
//		"csi-vol-95fcee27-e027-11ec-a257-764b17edad84",
//		"csi-vol-c3a6fa33-2fcd-11ed-b12f-8af21f51cb6d",
//		"csi-vol-150af83a-1f37-11ef-a388-c6f2daeeb83a",
//		"csi-vol-e80e1dcd-efd7-11ed-b280-9e559261c936",
//		"csi-vol-a0828496-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-82a501a3-0603-11ee-a9fd-3ebe5e805178",
//		"csi-vol-21d24bc1-2fb6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-f2b24ab8-efda-11ed-b280-9e559261c936",
//		"csi-vol-1d195d37-c7f1-11ed-9bbe-baccb3867649",
//		"csi-vol-7edd7e25-6429-11ee-9abc-ea01a90ca1db",
//		"csi-vol-9b381bd1-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-1cdf88e8-efdc-11ed-b280-9e559261c936",
//		"csi-vol-5eda40c7-2cfd-11ed-9914-428aff3909c1",
//		"csi-vol-eab2cd4c-1ea0-11ef-a388-c6f2daeeb83a",
//		"csi-vol-6dbd7223-2fc0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-c543f05f-2fcd-11ed-b12f-8af21f51cb6d",
//		"csi-vol-3edf34e5-efdc-11ed-b280-9e559261c936",
//		"csi-vol-e67d556e-4433-11ef-a388-c6f2daeeb83a",
//		"csi-vol-ceb6b365-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-ede38ae9-867f-4367-a2d4-17a7a084f180",
//		"csi-vol-b90a599b-5089-11ed-b12f-8af21f51cb6d",
//		"csi-vol-2eb22e6e-1a2e-11ed-b73f-7adb244cf5bc",
//		"csi-vol-77530aaf-adf5-11ed-b9fc-7a55e2a23421",
//		"csi-vol-dd564057-2fcf-11ed-b12f-8af21f51cb6d",
//		"csi-vol-e77c0087-ef13-11ed-b280-9e559261c936",
//		"csi-vol-7ffd788c-adf5-11ed-b9fc-7a55e2a23421",
//		"csi-vol-a65b7d9d-f735-469d-ba9c-1f9f7d701129",
//		"csi-vol-fc688c66-4957-11ed-ba52-c69a66c45312",
//		"csi-vol-65d466d7-e501-11ed-b7fb-da6452cd87b3",
//		"csi-vol-4781bda0-2fb6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-e8822e84-5630-11ef-9e71-66877a12684d",
//		"csi-vol-bf43efad-66a1-11ee-9abc-ea01a90ca1db",
//		"csi-vol-ff125f85-5daf-4ede-9a40-43b620c8b0d2",
//		"csi-vol-f5631b5a-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-e1c982f9-49be-11ef-a388-c6f2daeeb83a",
//		"csi-vol-9141900e-0d35-11ef-9823-36560c924df3",
//		"csi-vol-b62accdf-50b1-11ee-9abc-ea01a90ca1db",
//		"csi-vol-b3be5d3e-5c84-4866-ada0-6ad1dad985cb",
//		"csi-vol-63f1eb14-0d2e-11ef-9823-36560c924df3",
//		"csi-vol-d72d1760-03c4-11ee-9b40-caf255bd9603",
//		"csi-vol-329995e3-adf5-11ed-b9fc-7a55e2a23421",
//		"csi-vol-35c9168a-bc8d-11ec-8db4-5abd84c443b3",
//		"csi-vol-e8d6d9c1-1a41-11ed-b73f-7adb244cf5bc",
//		"csi-vol-6039c783-54a8-11ef-9e71-66877a12684d",
//		"csi-vol-956a848d-4438-11ef-a388-c6f2daeeb83a",
//		"csi-vol-e8d2940b-1a41-11ed-b73f-7adb244cf5bc",
//		"csi-vol-ef83ce17-2fce-11ed-b12f-8af21f51cb6d",
//		"csi-vol-a0d016d9-77d1-11ee-9abc-ea01a90ca1db",
//		"csi-vol-5f01aedc-2550-11ee-866b-2a544f3a69ac",
//		"csi-vol-76ea2368-efdc-11ed-b280-9e559261c936",
//		"csi-vol-e7cdc2dc-ef13-11ed-b280-9e559261c936",
//		"csi-vol-f2fb5b27-6929-11ef-9f88-e60b227d963c",
//		"csi-vol-ce9ab1b4-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-28cc242c-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-e36f6f76-9745-49fc-a1ca-1b3e45af0c61",
//		"csi-vol-2d166c83-7569-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-45b579c2-2fc6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-26e3d02e-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-a5c73dac-efdc-11ed-b280-9e559261c936",
//		"csi-vol-75d8c3a9-ca2b-11ed-9bbe-baccb3867649",
//		"csi-vol-d2684f08-a5d9-4134-865a-938f34c7f0fc",
//		"csi-vol-d1a1c9b0-bef2-11ec-8db4-5abd84c443b3",
//		"csi-vol-9b55b926-2fbe-11ef-a388-c6f2daeeb83a",
//		"csi-vol-890cd78c-2fa7-11ed-b12f-8af21f51cb6d",
//		"csi-vol-50ab2ee4-efd5-11ed-b280-9e559261c936",
//		"csi-vol-2d253adf-7569-11ed-97ed-9ea2adbd79a6",
//		"csi-vol-3ef44909-efdc-11ed-b280-9e559261c936",
//		"csi-vol-669e7634-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-e61a15dd-4fe6-11ef-a388-c6f2daeeb83a",
//		"csi-vol-1e32166b-f090-11ed-b280-9e559261c936",
//		"csi-vol-dd6666f2-2fcf-11ed-b12f-8af21f51cb6d",
//		"csi-vol-434f827e-326d-41c2-891a-4feb74e5c0ab",
//		"csi-vol-9f438137-efda-11ed-b280-9e559261c936",
//		"csi-vol-53cd9559-1f37-11ef-a388-c6f2daeeb83a",
//		"csi-vol-0b28968c-2fb5-11ed-b12f-8af21f51cb6d",
//		"csi-vol-12af09df-569e-11ed-b12f-8af21f51cb6d",
//		"csi-vol-71e7f1c6-2560-11ee-ae23-6e9adfdafb04",
//		"csi-vol-5063f35d-1b67-11ee-a9fd-3ebe5e805178",
//		"csi-vol-72a5dcc3-4070-11ef-a388-c6f2daeeb83a",
//		"csi-vol-e8dc948d-1a41-11ed-b73f-7adb244cf5bc",
//		"csi-vol-2b9bf777-0d44-11ef-823b-9ed519677c90",
//		"csi-vol-99d4be1a-7185-11ee-9b52-82b82dddcb3e",
//		"csi-vol-c0d9d7bc-2fcd-11ed-b12f-8af21f51cb6d",
//		"csi-vol-185ac814-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-9f575914-efda-11ed-b280-9e559261c936",
//		"csi-vol-7e036991-0d44-11ef-823b-9ed519677c90",
//		"csi-vol-74c5ad1f-a8ef-413b-9afe-c2142601591a",
//		"csi-vol-e77bf56d-ef13-11ed-b280-9e559261c936",
//		"csi-vol-8c5f1394-443e-11ef-a388-c6f2daeeb83a",
//		"csi-vol-a0db6fa8-77d1-11ee-9abc-ea01a90ca1db",
//		"csi-vol-4496e9d6-2fce-11ed-b12f-8af21f51cb6d",
//		"csi-vol-65610a37-02f6-11ef-a388-c6f2daeeb83a",
//		"csi-vol-2a1cf6d8-4420-11ef-a388-c6f2daeeb83a",
//		"csi-vol-30ca2e2c-69e3-11ef-9e43-ca882da5a549",
//		"csi-vol-b7eaa2ee-5fb0-11ef-9e71-66877a12684d",
//		"csi-vol-4946bcc2-2cf6-11ed-9914-428aff3909c1",
//		"csi-vol-b8f9b630-5089-11ed-b12f-8af21f51cb6d",
//		"csi-vol-a070ab7b-2fca-11ed-b12f-8af21f51cb6d",
//		"csi-vol-770a7e8b-e0cb-11ec-a257-764b17edad84",
//		"csi-vol-0403a31c-2d44-11ef-9f23-fe3532ece709",
//		"csi-vol-34c47851-2fd0-11ed-b12f-8af21f51cb6d",
//		"csi-vol-eab2935b-1ea0-11ef-a388-c6f2daeeb83a",
//		"csi-vol-2c7a195c-69e3-11ef-9e43-ca882da5a549",
//		"csi-vol-86860b4f-2fce-11ed-b12f-8af21f51cb6d",
//		"csi-vol-ca82e79e-efdc-11ed-b280-9e559261c936",
//		"csi-vol-6f431b92-4da4-11ef-a388-c6f2daeeb83a",
//		"csi-vol-2554ba12-4064-11ef-a388-c6f2daeeb83a",
//		"csi-vol-e9cdde05-441f-11ef-a388-c6f2daeeb83a",
//		"csi-vol-0d4bfbdf-9d87-11ee-a30a-92d7948fed0a",
//		"csi-vol-87783796-2fb4-11ed-b12f-8af21f51cb6d",
//		"csi-vol-bbe3af1f-3484-11ef-a388-c6f2daeeb83a",
//		"csi-vol-d0ac658f-381e-400f-8614-8e3821e25566",
//		"csi-vol-d9c0bc13-2fce-11ed-b12f-8af21f51cb6d",
//		"csi-vol-35668fda-0602-11ee-a9fd-3ebe5e805178",
//		"csi-vol-04039487-2d44-11ef-9f23-fe3532ece709",
//		"csi-vol-8c5cee67-443e-11ef-a388-c6f2daeeb83a",
//		"csi-vol-c0efe173-2fcd-11ed-b12f-8af21f51cb6d",
//		"csi-vol-95fd8659-e027-11ec-a257-764b17edad84",
//		"csi-vol-3e13f691-78b5-4dbe-8d56-7c46b8ed796e",
//		"csi-vol-574078cc-f938-11ed-9b40-caf255bd9603",
//		"csi-vol-1632d6cc-443d-11ef-a388-c6f2daeeb83a",
//		"csi-vol-770d0032-e0cb-11ec-a257-764b17edad84",
//		"csi-vol-d1ac0a28-2fc8-11ed-b12f-8af21f51cb6d",
//		"csi-vol-45badf55-2fc6-11ed-b12f-8af21f51cb6d",
//		"csi-vol-5f062013-2550-11ee-866b-2a544f3a69ac",
//		"csi-vol-b618e791-50b1-11ee-9abc-ea01a90ca1db",
//	}
var subvolumeNames []string

var pvFromKubeArr []string

var LOG_LEVEL = getEnv("LOG_LEVEL", "INFO")

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

	// Загрузка kubeconfig
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

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
	logWithTime(fmt.Sprintf("🌽Matches: %+v", matches))
	logWithTime(fmt.Sprintf("🦐Differences: %+v", differences))
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
