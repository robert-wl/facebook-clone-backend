package database

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

const timeLayout = "2006-01-02 15:04:05"

func generateImage(resolution *[]string) string {
	images := []string{
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239889599-zeubp5w4zd-7afb0491c91b2f9e9aac56667c3be677.jpg?alt=media\u0026token=611497b0-4729-4a3c-a712-e5b030c76c98",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239889927-binnrjn0rls-Wonders-of-the-World-Pyramids-1030x538.png?alt=media\u0026token=60946674-0a1f-450f-88f6-3cdd452005e8",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239890364-4lgu1idb3f3-header.jpg?alt=media\u0026token=67cbc129-8050-4de2-b81d-632a00849b54",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239890642-bgv6gy7619g-wp4535284.webp?alt=media\u0026token=ae06e0fd-aa17-432c-badf-4cbe8022db30",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891088-2q3ng215db7-pixel-art-creature-sword-hyper-light-drifter-wallpaper-preview.jpg?alt=media\u0026token=9c3dfa20-6eec-4307-8a71-0ead6ff977cd",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891289-s5iziys6x8-aesthetic-super-mario-running-desktop-wallpaper-preview.jpg?alt=media\u0026token=d49fd349-994a-414e-8b46-d7de4a33e11d",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891629-1nprmy82w9j-elden-ring-landscape-game-art-video-game-art-video-games-hd-wallpaper-preview.jpg?alt=media\u0026token=4b4b01c3-f59e-4f9c-9914-43d320393a86",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891872-f2bskviqt7p-HD-wallpaper-anime-landscape-ai-city.jpg?alt=media\u0026token=2cf51657-d5d0-40a5-b2fe-1b5fc1f41b30\",\"directory\": \"post/1722239891872-f2bskviqt7p-HD-wallpaper-anime-landscape-ai-city.jpg",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239892077-4ylhf2mrqdt-anime-landscape-anime-art-painting-sea-wallpaper-preview.jpg?alt=media\u0026token=b35c31fb-a4e8-4ca5-8345-a245c6f1c43d",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239892440-bvg0j89wa9-88016860_p0_master1200.jpg?alt=media\u0026token=59dd9e3a-1d3f-43e3-b6e8-0eea996171b7",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722238113910-m8xehm01d1-image_5.png?alt=media\u0026token=c3a85375-58f6-4636-b435-a7fc04969b9b",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722238114728-depr6qvc1z7-image_6.png?alt=media\u0026token=7408b912-d806-4995-b05f-5006745f426e",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722238115256-9tfeismfm7w-88016860_p0_master1200.jpg?alt=media\u0026token=8256a127-59cc-4caf-8830-a42f94d66c41",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/story%2F1691605952816-q8j4hn0vz6-93960519_p0_master1200.jpg?alt=media&token=7d849d47-312b-45e8-8da0-c4e4a24ba7aa",
	}

	resolutions := [][]string{
		{"1920", "1080"},
		{"1280", "720"},
		{"1366", "768"},
		{"1600", "900"},
		{"800", "600"},
		{"1024", "768"},
		{"1280", "1024"},
		{"720", "480"},
	}

	if rand.Intn(10) > 8 {
		image := images[rand.Intn(len(images))]
		data := fmt.Sprintf("{\"url\": \"%s\",\"directory\": \"%s\",\"type\": \"image/jpeg\"}", image, image)

		return data
	} else {
		var currRes []string
		if resolution == nil {
			currRes = resolutions[rand.Intn(len(resolutions))]
		} else {
			currRes = *resolution
		}

		id := rand.Intn(1000)

		image := fmt.Sprintf("https://picsum.photos/id/%d/%s/%s", id, currRes[0], currRes[1])
		data := fmt.Sprintf("{\"url\": \"%s\",\"directory\": \"%s\",\"type\": \"image/jpeg\"}", image, image)

		return data
	}
}

func generateProfile() string {
	return fmt.Sprintf("https://i.pravatar.cc/300?img=%d", rand.Intn(70)+1)
}

func generateVideo() string {
	videos := []string{
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251656896-7ipol8g19se-Y2meta.app-CURIOSITY%20-%20Featuring%20Richard%20Feynman-(1080p).mp4?alt=media\u0026token=27ba7263-6665-4adc-807a-4c1838781faa",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251660098-2b4lh2hkqd1-Y2meta.app-ocean-(1080p).mp4?alt=media\u0026token=13589535-70bb-434e-ab90-aece2b8f6162",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251681222-59123q78wri-Y2meta.app-Indonesia%20_%20Cinematic%20Travel%20Video%20_%20Stock%20Footage-(1080p).mp4?alt=media\u0026token=49353b71-e427-41e5-859e-49e6293b0120",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251717038-b6zphnp94sf-Y2meta.app-JAPAN%20-%20Like%20You've%20Never%20Seen%20Before%20_%20Stock%20Footage-(1080p).mp4?alt=media\u0026token=f5332066-844c-49b3-8ed5-2a5eb37d54ac",
	}

	video := videos[rand.Intn(len(videos))]
	data := fmt.Sprintf("{\"url\": \"%s\",\"directory\": \"%s\",\"type\": \"video/mp4\"}", video, video)

	return data
}

func generateReelVideo() string {
	videos := []string{
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474485622-bneunn4a7vu-Our%20brain%20ain%E2%80%99t%20braining%20anymore%20%20.Tag%20your%20best%20friends%20down%20below%20if%20it%E2%80%99s%20relatable%20.Follow%20%40phrivn%20for%20more%20daily%20content%20%20%20..%23fashion%20%23ootd%20%23reels%20%20%20menswear%2C%20men%E2%80%99s%20fashion%2C%20fashion%20inspo%2C%20fashion%20inspiration%2C.mp4?alt=media&token=80d19253-67c4-4269-9c09-dc32bd27a57b",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507617794-l32boymopa9-by%20%40hamigua_bifido%20%20Ever%20heard%20about%20the%20duck%20that%20loved%20to%20wear%20hats%20%20One%20day%2C%20a%20duck%20named%20Quacky%20discovered%20a%20basket%20of%20tiny%20hats%20left%20at%20a%20picnic.%20%20Quacky%20started%20trying%20them%20on%2C%20waddling%20around%20in%20different%20s.mp4?alt=media&token=b3cb6cdf-eaa1-4c23-b60a-15384d164320",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474240350-jgccsz7qr3-Sure%20thing!%20Here%E2%80%99s%20the%20statistics%20about%20the%20Honda%20Civic%20Type%20R-The%20Honda%20Civic%20Type%20R%20is%20a%20remarkable%20performance%20car%20renowned%20for%20its%20agile%20handling%20and%20sporty%20design.%20Equipped%20with%20a%20robust%202.0-liter%20turbocharge.mp4?alt=media&token=12fda4da-fcf4-43f9-a508-2ac2b0422df0",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474507822-j6ek2cgeco-Sure%20thing!%20Here%E2%80%99s%20the%20statistics%20about%20the%20Honda%20Civic%20Type%20R-The%20Honda%20Civic%20Type%20R%20is%20a%20remarkable%20performance%20car%20renowned%20for%20its%20agile%20handling%20and%20sporty%20design.%20Equipped%20with%20a%20robust%202.0-liter%20turboch%20(1).mp4?alt=media&token=f404a28a-1915-4489-a3cb-6f24258109d1",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507720517-vutblib5ao-%E2%80%9CGet%20a%20load%20of%20these%20guys%E2%80%9D%23dog%20%23doglover%20%23dogsofinstagram.mp4?alt=media&token=5673cd2b-5af7-4191-90c5-9913d4169675",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507665435-wmopmjhqk9g-Day%20191.mp4?alt=media&token=67ba516d-4076-4303-bdab-01111dc51b47",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474409926-qb9qx06vngs-That%E2%80%99s%20crazy%20%20%20%23olympics%20%23paris%20%23eiffeltower%20%23olympics2024%20%23TeamUSA%20%E2%80%A2%E2%80%A2%E2%80%A2%E2%80%A2%20%23travelgram%20%23exploremore%20%23wanderlust%20%23viralreel%20%23explorepage%20%23trending%23instaviral%20%23reelsinstagram%20%23viralvideos%23fyp%20%23viral%20%23reelitfeelit%20%23wan.mp4?alt=media&token=c83ce1d2-8314-4622-83b0-fbaff1372da8",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507677924-u05uaqitcoe-%EC%82%AC%EB%82%98%EC%9A%B4%20%EB%9D%BC%EC%BF%A4%20%EB%B9%A8%EB%9E%98%ED%95%98%EB%8A%94%EB%B0%A9%EB%B2%95...%23%EA%B1%B4%EB%8C%80%20%23%EB%9D%BC%EC%BF%A4%20%23%EB%AF%B8%EC%96%B4%EC%BA%A3%EC%A1%B1%EC%9E%A5%20%23%EA%B1%B4%EB%8C%80%EC%B9%B4%ED%8E%98%20%23%EA%B1%B4%EB%8C%80%ED%95%AB%ED%94%8C%20%23%EA%B1%B4%EB%8C%80%EB%8D%B0%EC%9D%B4%ED%8A%B8%20%23%EA%B1%B4%EB%8C%80%EB%A7%9B%EC%A7%91%20%23%EC%84%B1%EC%88%98%EB%8F%99%20%23%EC%84%B1%EC%88%98%EB%8F%99%EC%B9%B4%ED%8E%98%20%23%EC%84%B1%EC%88%98%EB%8F%99%EB%A7%9B%EC%A7%91%20%23%EC%8B%A4%EB%82%B4%EB%8D%B0%EC%9D%B4%ED%8A%B8%20%23%EB%8D%B0%EC%9D%B4%ED%8A%B8%20%23%EB%8D%B0%EC%9D%B4%ED%8A%B8%EC%BD%94%EC%8A%A4%20%23seoul%20%23seoulkorea%20%23raccoon.mp4?alt=media&token=b6c7d925-04f0-485c-a52a-0ee8e5797b2c",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507272488-p25kixuf1sf-Here%E2%80%99s%20the%20information%20about%20the%20Mercedes%20CLR%20GTR-The%20Mercedes%20CLR%20GTR%20is%20a%20remarkable%20racing%20car%20celebrated%20for%20its%20outstanding%20performance%20and%20sleek%20design.%20Powered%20by%20a%20potent%206.0-liter%20V12%20engine%2C%20it%20delivers%20.mp4?alt=media&token=0a40a868-84a0-4661-92b0-66f6df2508f4",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507648274-fno530ei7nd-Rolls-Royce%20is%20synonymous%20with%20luxury%2C%20elegance%2C%20and%20exceptional%20craftsmanship%2C%20setting%20the%20standard%20in%20the%20ultra-luxury%20automotive%20segment.%20Known%20for%20its%20opulent%20design%20and%20meticulous%20attention%20to%20detail%2C%20Rolls-R.mp4?alt=media&token=a9ba9810-58ae-4b95-9e39-875a1b4692be",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507657148-exwj588mjkl-Day%20218%23sneezing%20%23cat%20%23daily%20%23meme%20The%20Tesla%20Cybertruck%20is%20an%20all-electric%2C%20battery-powered%20light-duty%20truck%20unveiled%20by%20Tesla%2C%20Inc.Here's%20a%20comprehensive%20overview%20of%20its%20key%20features%20and%20specifications-Tesla%20Cybe.mp4?alt=media&token=5c6d1af4-9344-4ec3-b784-39069e91c5a2",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507769411-1b0hoq77rtx-finding-breakers-fuses-electronic-technology-funny-experiment-720-ytshorts.savetube.me.mp4?alt=media&token=3abda830-1f2d-4825-92de-b18ac800702d",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507691773-522yu3ucnih-Very%20round%20chicken%20%23Magic%20%23Animal%20%23Chicken%20%23Share.mp4?alt=media&token=299d639d-e558-4619-9760-822fc492ff12",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474455071-4125ptzq4zk-Not%20their%20first%20potato%20sack%20race%20%20%20(via%20dark_emperor16-TT).mp4?alt=media&token=632a14c4-6850-4c08-9ca4-59e1cdde3ff1",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474493550-6d4lo77opmc-Can%20a%20Seed%20Grow%20in%20Your%20Nose.mp4?alt=media&token=849f9b9b-80aa-4a37-a23d-a32e5599f077",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507764084-bv7rn34q3lm-i-m-still-astounded-this-is-true-720-ytshorts.savetube.me.mp4?alt=media&token=e7dd5686-e33d-4887-8d76-631debe1cc15",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507758679-o2ov6nn79xm-guess-the-minecraft-block-in-60-seconds-42-720-ytshorts.savetube.me.mp4?alt=media&token=c360bb88-0ba0-479e-bf75-6a23c62a3748",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507671474-nl5b2vf418o-classicString%20theory%20is%20a%20theoretical%20framework%20in%20physics%20that%20aims%20to%20reconcile%20general%20relativity%20and%20quantum%20mechanics%20%20.%20It%20proposes%20that%20the%20fundamental%20particles%20we%20observe%20are%20not%20point-like%20dots%2C%20but%20rath.mp4?alt=media&token=f7a5275f-30ac-43ac-91c4-5d91da4af568",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507635354-wkujxfn3fxp-Normal%20physiological%20levels%20of%20plasma%20C-peptide%20during%20fasting%20range%20from%200.9%20-%201.8%20ng-ml.%20Higher%20levels%20may%20suggest%20conditions%20such%20as%20insulin%20resistance%2C%20insulinoma%2C%20or%20kidney%20disease.%20On%20the%20other%20hand%2C%20reduced.mp4?alt=media&token=399e791e-a5cd-4da5-8dbd-c724ebda3441",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474471126-vcwg6wmpdf-The%20Tesla%20Cybertruck%20is%20an%20all-electric%2C%20battery-powered%20light-duty%20truck%20unveiled%20by%20Tesla%2C%20Inc.Here's%20a%20comprehensive%20overview%20of%20its%20key%20features%20and%20specifications-Tesla%20Cybertruck%20OverviewDesign%20and%20Structure.mp4?alt=media&token=1072dae4-244a-42b6-b1ee-4202c58064a8",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507734893-x8475s8j0zj-Follow%20me%20%20%20and%20The%20Tesla%20Cybertruck%20is%20an%20all-electric%2C%20battery-powered%20light-duty%20truck%20unveiled%20by%20Tesla%2C%20Inc.Here%E2%80%99s%20a%20comprehensive%20overview%20of%20its%20key%20features%20and%20specifications-Tesla%20Cybertruck%20OverviewDesi.mp4?alt=media&token=67f076d0-b524-4d1a-b10e-d3ae992fa9ea",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474462606-tw1n4m2u5to-Getting%20Struck%20By%20Lightening%20Twice.mp4?alt=media&token=f0ff4b2b-9080-4826-b2cc-fc367fd60856",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474513854-7trigmbjpjj-Dear%20friend%2C%20I%E2%80%99m%20coming%20%20%20%23fy%20%23fyp%20%23trending%20%23cute%20%23adorable%20%23redpanda%20%23love.mp4?alt=media&token=d80e8d81-0d68-4a61-90f3-3f9703af6319",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507259967-lbcyxh5v0oh-'Show%20me%20your%20motivation'%23meme%23memes%23memesdaily%23memesdaily%23funnymemes%23funnymemes%23funnymemesdaily%23shitpost%23shitposts%23shitposting%23shitpostmemes%23schizoposting%23schizomemes%23schizophrenia%23eldenring%23eldenringmemes.mp4?alt=media&token=865f732c-2729-4d7e-abd1-b10f76c98bea",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507752995-n5bt8uadcsf-Here%E2%80%99s%20the%20information%20about%20the%20Mercedes%20CLR%20GTR-The%20Mercedes%20CLR%20GTR%20is%20a%20remarkable%20racing%20car%20celebrated%20for%20its%20outstanding%20performance%20and%20sleek%20design.%20Powered%20by%20a%20potent%206.0-liter%20V12%20engine%2C%20it%20deliv%20(1).mp4?alt=media&token=4e0cedc7-8ba4-4d1d-a90f-e39dcf3a2e4b",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507684867-gm3gfe0vpbh-An-o-n51iQEBoU8Knc49n_rRoW479zyk9-Bf9KaXgkCNi7Z6d6ncTNNZpP-s-Ss-5l2WxD2hkGJi17z4a7F23eA.mp4?alt=media&token=16ba65d6-c4e1-4812-a3e7-8f0d81ec4402",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507628115-kvu719quh2p-song-%20%E2%80%9Ca%20safe%20place%20to%20stay%E2%80%9D%20%20%F0%9F%AA%BB%20%23catvideos%20%23quotes%20%23motivation%20%23hopecore%20%23positivity%20%23darkambient....Producer%20music%20dark%20ambient%20quotes%20motivation%20hopecore%20musician%20funny%20cat%20videos.mp4?alt=media&token=434401e1-1535-4365-81f7-4cff4c772bb6",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507745427-8rsm08d39i-The%20Tesla%20Cybertruck%20is%20an%20all-electric%2C%20battery-powered%20light-duty%20truck%20unveiled%20by%20Tesla%2C%20Inc.Here's%20a%20comprehensive%20overview%20of%20its%20key%20features%20and%20specifications-Tesla%20Cybertruck%20OverviewDesign%20and%20Struc%20(1).mp4?alt=media&token=6bf76ebb-d16e-4485-ab02-34030f453a2d",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722507090517-uyp0oirg0zb-instagram%20-%203419300499168940664%20-%20hop_in_games%20-%20C9zy7ARR9J4.mp4?alt=media&token=5e93e264-5b0c-4e00-9ab1-cd4d581a1987",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722474501342-tgxdvaid9fa-Bro%20was%20prepared%20%20Watch%20as%20this%20pint-sized%20adventurer%2C%20the%20small%20kitten%2C%20takes%20on%20the%20monumental%20challenge%20of%20climbing%20the%20stairway%2C%20one%20tiny%20step%20at%20a%20time.%20%20%20But%20wait%2C%20there%E2%80%99s%20a%20heartwarming%20twist%20%E2%80%93%20a%20big-hearte.mp4?alt=media&token=d15baee5-6397-4db6-bd8b-eee54293bf43",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/reels%2F1722444447811-ofyop9iwbis-Snapinsta.app_video_An9uSGOY9L75rgwBruOOMOLZ_fdYmxYmlNemHV59yDQqSab7X4Kmr-WlobxyYiKxvZIj75ZJiADLZNniCacRLDV_.mp4?alt=media&token=9cbecde8-4288-4292-9de0-428650990cfa",
	}

	video := videos[rand.Intn(len(videos))]

	return video
}

func generateGroupFile() []string {
	fileData := [][]string{
		{"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/groups%2F1722690627861-d8ec04qgk4p-shinchan.gif?alt=media&token=4517df33-5dba-4468-ab24-eb7fa53019b0",
			"shinchan.gif",
			"image/gif"},
		{"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/groups%2F1722690571678-3ax7t6ajm6g-github.txt?alt=media&token=ea36ecb9-94e1-435d-a8ac-ef01f7d23553",
			"github.txt",
			"text/plain"},
		{"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/groups%2F1722690762111-ufragftou4h-7%20to%203.mp3?alt=media&token=1c2ae371-9435-49cd-ae1a-c5f48b3af5dd",
			"7 to 3.mp3",
			"audio/mpeg"},
		{"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/groups%2F1722690693839-6xinic1u4v5-generic-table.tsx?alt=media&token=20b0dccf-e532-4a21-b44b-3acdbbf21692",
			"generic-table.tsx",
			"text/plain"},
		{"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/groups%2F1722690784565-tc3aa14kw4o-April%20Showers.mp3?alt=media&token=fd0ec80b-9e89-46ed-82e0-821c84517206",
			"April Showers.mp3",
			"audio/mpeg"},
		{"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/groups%2F1722690354148-umexa9uw8o9-hld.png?alt=media&token=e4434cd0-db98-4143-a63c-0d17c5aa6412",
			"hld.png",
			"image/png"},
		{"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/groups%2F1722690232438-1z2b73zv9ja-hello.txt?alt=media&token=b6fa74be-546b-4d9b-b254-2018437b9b0d",
			"hello.txt",
			"text/plain"},
	}

	return fileData[rand.Intn(len(fileData))]
}

func generateUser() []model.User {
	var users []model.User

	db := GetDBInstance()
	for i := 0; i < 40; i++ {
		fmt.Println("Generating User")

		pw, _ := utils.EncryptPassword("password")

		fn := faker.FirstName()
		ln := faker.LastName()

		profile := generateProfile()
		background := generateImage(&([]string{"1920", "1080"}))

		var gender string

		if rand.Intn(10) > 5 {
			gender = "Male"
		} else {
			gender = "Female"
		}

		dob := time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000)))
		user := model.User{
			ID:                uuid.NewString(),
			FirstName:         fn,
			LastName:          ln,
			Username:          fmt.Sprintf("%s%s", fn, ln),
			Email:             fmt.Sprintf("%s.%s@gmail.com", fn, ln),
			Password:          pw,
			Dob:               dob,
			Gender:            gender,
			Active:            true,
			MiscId:            nil,
			Profile:           &profile,
			Background:        &background,
			CreatedAt:         time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
			FriendCount:       0,
			MutualCount:       0,
			NotificationCount: 0,
			Friended:          "",
			Theme:             "light",
		}

		db.Create(&user)

		users = append(users, user)
	}

	return users
}

func generatePost(user model.User) []model.Post {
	var posts []model.Post
	db := GetDBInstance()

	num := rand.Intn(20) + 1
	for i := 0; i < num; i++ {
		fmt.Println("Generating Post")

		var files []*string
		if rand.Intn(10) > 4 {
			if rand.Intn(10) > 8 {
				data := generateVideo()
				files = append(files, &data)

			} else {
				takeAmount := rand.Intn(10) + 1

				for i := 0; i < takeAmount; i++ {
					image := generateImage(nil)
					files = append(files, &image)
				}
			}
		}

		post := model.Post{
			ID:         uuid.NewString(),
			UserID:     user.ID,
			Content:    faker.Sentence(),
			Privacy:    "public",
			ShareCount: rand.Intn(100),
			CreatedAt:  time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
			Files:      files,
		}

		db.Create(&post)

		posts = append(posts, post)
	}

	return posts
}

func generateFriend(users []model.User) {
	db := GetDBInstance()

	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			fmt.Println("Generating Friend")
			sender := users[i]
			receiver := users[j]

			friend := model.Friend{
				SenderID:   sender.ID,
				ReceiverID: receiver.ID,
				Accepted:   true,
			}
			friend2 := model.Friend{
				SenderID:   receiver.ID,
				ReceiverID: sender.ID,
				Accepted:   true,
			}

			db.Create(&friend)
			db.Create(&friend2)
		}
	}
}

func generatePostLike(users []model.User, posts []model.Post) {
	db := GetDBInstance()

	for _, user := range users {
		for _, post := range posts {
			if rand.Intn(10) > 4 {
				fmt.Println("Generating Post Like")
				postLike := model.PostLike{
					UserID: user.ID,
					PostID: post.ID,
				}

				db.Create(&postLike)
			}
		}
	}
}

func generatePostComment(users []model.User, posts []model.Post) []model.Comment {
	db := GetDBInstance()

	var comments []model.Comment
	for _, user := range users {
		for _, post := range posts {
			if rand.Intn(10) > 8 {
				fmt.Println("Generating Post Comment")
				postComment := model.Comment{
					ID:           uuid.NewString(),
					UserID:       user.ID,
					Content:      faker.Sentence(),
					ParentPostID: &post.ID,
					CreatedAt:    time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
				}

				db.Create(&postComment)
				comments = append(comments, postComment)
			}
		}
	}

	for _, comment := range comments {
		for _, user := range users {
			if rand.Intn(10) > 8 {
				fmt.Println("Generating Comment Reply")
				commentReply := model.Comment{
					ID:              uuid.NewString(),
					UserID:          user.ID,
					Content:         faker.Sentence(),
					ParentCommentID: &comment.ID,
					CreatedAt:       time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
				}

				db.Create(&commentReply)
				comments = append(comments, commentReply)
			}
		}
	}

	return comments
}

func generateCommentLike(users []model.User, comments []model.Comment) {
	db := GetDBInstance()

	for _, user := range users {
		for _, comment := range comments {
			if rand.Intn(10) > 7 {
				fmt.Println("Generating Comment Like")
				commentLike := model.CommentLike{
					UserID:    user.ID,
					CommentID: comment.ID,
				}

				db.Create(&commentLike)
			}
		}
	}

}

func generateConversation(users []model.User) {
	db := GetDBInstance()

	randUser := users
	for i := 0; i < len(users); i++ {
		rand.Shuffle(len(randUser), func(i, j int) {
			randUser[i], randUser[j] = randUser[j], randUser[i]
		})

		randLength := rand.Intn(len(users))

		for j := 0; j < randLength; j++ {
			if randUser[j].ID == randUser[i].ID {
				continue
			}

			fmt.Println("Generating Conversation")
			sender := randUser[i]
			receiver := randUser[j]

			conversation := model.Conversation{
				ID: uuid.NewString(),
			}

			db.Create(&conversation)

			cUser1 := model.ConversationUsers{
				ConversationID: conversation.ID,
				UserID:         sender.ID,
			}

			cUser2 := model.ConversationUsers{
				ConversationID: conversation.ID,
				UserID:         receiver.ID,
			}

			db.Create(&cUser1)
			db.Create(&cUser2)
		}
	}
}

func generateStories(users []model.User) {
	db := GetDBInstance()

	colors := []string{
		"lightblue",
		"pink",
		"lightgray",
		"orange",
	}

	fonts := []string{
		"normal",
		"roman",
	}

	for _, user := range users {
		for i := 0; i < rand.Intn(10); i++ {
			fmt.Println("Generating Story")
			textBr := faker.Sentence()

			if rand.Intn(10) > 5 {
				story := model.Story{
					ID:        uuid.NewString(),
					UserID:    user.ID,
					Font:      &fonts[rand.Intn(len(fonts))],
					Color:     &colors[rand.Intn(len(colors))],
					Text:      &textBr,
					CreatedAt: time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(20))),
				}

				db.Create(&story)
			} else {
				image := generateImage(nil)
				story := model.Story{
					ID:        uuid.NewString(),
					UserID:    user.ID,
					Image:     &image,
					CreatedAt: time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(20))),
				}

				db.Create(&story)
			}
		}
	}
}

func generateReels(users []model.User) []model.Reel {
	db := GetDBInstance()

	reels := []model.Reel{}
	for _, user := range users {
		for i := 0; i < rand.Intn(10); i++ {
			fmt.Println("Generating Reel")
			video := generateReelVideo()

			reel := model.Reel{
				ID:         uuid.NewString(),
				UserID:     user.ID,
				Content:    faker.Sentence(),
				Video:      video,
				ShareCount: rand.Intn(100),
				CreatedAt:  time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
			}

			db.Create(&reel)

			reels = append(reels, reel)
		}
	}

	return reels
}

func generateReelComment(users []model.User, reels []model.Reel) []model.ReelComment {
	db := GetDBInstance()

	comments := []model.ReelComment{}
	for _, user := range users {
		for _, reel := range reels {
			if rand.Intn(10) > 8 {
				fmt.Println("Generating Reel Comment")
				comment := model.ReelComment{
					ID:           uuid.NewString(),
					UserID:       user.ID,
					Content:      faker.Sentence(),
					ParentReelID: &reel.ID,
					CreatedAt:    time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
				}

				db.Create(&comment)
				comments = append(comments, comment)
			}
		}
	}

	for _, comment := range comments {
		for _, user := range users {
			if rand.Intn(10) > 8 {
				fmt.Println("Generating Reel Comment Reply")
				commentReply := model.ReelComment{
					ID:              uuid.NewString(),
					UserID:          user.ID,
					Content:         faker.Sentence(),
					ParentCommentID: &comment.ID,
					CreatedAt:       time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
				}

				db.Create(&commentReply)
				comments = append(comments, commentReply)
			}
		}
	}

	return comments

}

func generateReelLike(users []model.User, reels []model.Reel) {
	db := GetDBInstance()

	for _, user := range users {
		for _, reel := range reels {
			if rand.Intn(10) > 7 {
				fmt.Println("Generating Reel Like")
				reelLike := model.ReelLike{
					UserID: user.ID,
					ReelID: reel.ID,
				}

				db.Create(&reelLike)
			}
		}
	}
}

func generateReelCommentLike(users []model.User, comments []model.ReelComment) {
	db := GetDBInstance()

	for _, user := range users {
		for _, comment := range comments {
			if rand.Intn(10) > 7 {
				fmt.Println("Generating Reel Comment Like")
				commentLike := model.ReelCommentLike{
					ReelCommentID: comment.ID,
					UserID:        user.ID,
				}

				db.Create(&commentLike)
			}
		}
	}
}

func generateGroup(users []model.User) map[string][]model.Member {
	db := GetDBInstance()
	var groupMemberMap = make(map[string][]model.Member)

	groupNum := rand.Intn(60) + 20

	for i := 0; i < groupNum; i++ {
		fmt.Println("Generating Group")

		var privacy string

		if rand.Intn(10) > 6 {
			privacy = "Public"
		} else {
			privacy = "Private"
		}

		group := model.Group{
			ID:          uuid.NewString(),
			Name:        faker.Sentence(),
			About:       faker.Paragraph(),
			Privacy:     privacy,
			Background:  generateImage(&[]string{"1920", "1080"}),
			MemberCount: 0,
			ChatID:      nil,
			Chat:        nil,
			CreatedAt:   time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
		}

		db.Create(&group)

		var groupUsers []model.Member
		for _, user := range users {
			if rand.Intn(10) > 8 {
				fmt.Println("Generating Group User")

				var role string

				if rand.Intn(10) > 8 {
					role = "Admin"
				} else {
					role = "member"
				}

				requested := false
				if role == "member" {
					requested = true
				}

				groupUser := model.Member{
					GroupID:   group.ID,
					UserID:    user.ID,
					Requested: requested,
					Approved:  true,
					Role:      role,
				}

				groupUsers = append(groupUsers, groupUser)

				db.Create(&groupUser)
			}
		}

		if len(groupUsers) == 0 {
			groupUser := model.Member{
				GroupID:   group.ID,
				UserID:    users[rand.Intn(len(users))].ID,
				Requested: false,
				Approved:  true,
				Role:      "Admin",
			}

			groupUsers = append(groupUsers, groupUser)
			db.Create(&groupUser)
		}

		groupMemberMap[group.ID] = groupUsers

		group.MemberCount = len(groupUsers)

		db.Save(&group)
	}

	return groupMemberMap
}

func generateGroupConversation(groups map[string][]model.Member) {
	db := GetDBInstance()

	for _, members := range groups {

		conversation := model.Conversation{
			ID:      uuid.NewString(),
			GroupID: &members[0].GroupID,
		}

		db.Create(&conversation)

		var group model.Group

		if err := db.Where("id = ?", members[0].GroupID).First(&group).Error; err != nil {
			continue
		}

		group.ChatID = &conversation.ID

		db.Save(&group)

		for i := 0; i < len(members); i++ {
			fmt.Println("Generating Group Conversation")

			conversation := model.ConversationUsers{
				UserID:         members[i].UserID,
				ConversationID: conversation.ID,
			}

			db.Create(&conversation)
		}
	}
}

func generateGroupPosts(groups map[string][]model.Member) {
	db := GetDBInstance()

	for _, members := range groups {
		for i := 0; i < rand.Intn(20); i++ {
			fmt.Println("Generating Group Post")

			var files []*string
			if rand.Intn(10) > 4 {
				if rand.Intn(10) > 8 {
					data := generateVideo()
					files = append(files, &data)

				} else {
					takeAmount := rand.Intn(10) + 1

					for i := 0; i < takeAmount; i++ {
						image := generateImage(nil)
						files = append(files, &image)
					}
				}
			}

			post := model.Post{
				ID:         uuid.NewString(),
				UserID:     members[rand.Intn(len(members))].UserID,
				Content:    faker.Sentence(),
				Privacy:    "group",
				ShareCount: rand.Intn(100),
				CreatedAt:  time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
				Files:      files,
				GroupID:    &members[0].GroupID,
			}

			db.Create(&post)
		}
	}
}

func generateGroupFiles(groups map[string][]model.Member) {
	db := GetDBInstance()

	for _, members := range groups {
		for i := 0; i < rand.Intn(200); i++ {
			if rand.Intn(10) > 8 {
				fmt.Println("Generating Group File")

				fileData := generateGroupFile()

				url := fmt.Sprintf("{\"url\": \"%s\",\"directory\": \"%s\",\"type\": \"%s\"}", fileData[0], fileData[1], fileData[2])
				file := model.GroupFile{
					ID:      uuid.NewString(),
					UserID:  members[rand.Intn(len(members))].UserID,
					GroupID: members[0].GroupID,
					Name:    fileData[1],
					Type:    fileData[2],
					URL:     url,
				}

				db.Create(&file)
			}
		}
	}
}

func FakeData() {
	users := generateUser()

	var posts []model.Post
	for _, user := range users {
		userPosts := generatePost(user)

		posts = append(posts, userPosts...)
	}

	generatePostLike(users, posts)

	comments := generatePostComment(users, posts)

	generateCommentLike(users, comments)

	generateFriend(users)

	generateConversation(users)

	generateStories(users)

	reels := generateReels(users)

	generateReelLike(users, reels)

	reelComments := generateReelComment(users, reels)

	generateReelCommentLike(users, reelComments)

	groupData := generateGroup(users)

	generateGroupConversation(groupData)

	generateGroupPosts(groupData)

	generateGroupFiles(groupData)

	fmt.Println("Data Generated")
}
