package realmdefense

const (
	// ApiMessageList is the API to list the in-game messaging system
	ApiMessageList = "/message/list"
	// ApiTime returns the server time, so that the game can check your device's clock is correct
	ApiTime = "/time"
	// ApiSave lets you do a "cloud save" of your in-game progress
	ApiSave = "/save"
	// ApiRecover lets you do a "cloud restore" of your in-game progress
	ApiRecover = "/recover"
	// ApiLoadSave compares the existing save with the server one
	ApiLoadSave = "/loadsave"
	// ApiTournamentQuery lets you query the tournament scores
	ApiTournamentQuery = "/tournament/query"
	// ApiTournamentScore lets you send your tournament score
	ApiTournamentScore = "/tournament/score"
	// ApiStaticTnNews gives you the tournament news
	ApiStaticTnNews = "/static/tn/news"
)

const (
	RealmAndroid = "Android"
	// TODO: realm iOS?
)

type LoadSaveRequest struct {
	Data string
	Id   string
	Seq  int
}

type LoadSaveResponse struct {
	LoadOrSave  bool
	Seq         int
	Data        string
	CountryCode string
}

type MessageListRequest struct {
	UserId string
}

type MessageListResponse struct {
	News []MessageListNews
}

type MessageListNews struct {
	Time    uint64
	Exp     uint64
	Data    string
	Regions string
}

type MessageData struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	From        string `json:"from"`
	Exp         int    `json:"exp"`
	Attachments string `json:"attachments"`
}

type TimeRequest struct {
}

type TimeResponse struct {
	Now uint64
}

type TournamentQuery struct {
	Realm string
	Uid   string
}

type TournamentResponse struct {
	Now        uint64
	Tid        string
	Status     int
	Data       string
	League     int
	NextLeague int
	Banned     bool
	History    []TournamentHistoryEntry
	Scores     []TournamentScore
}

type TournamentHistoryEntry struct {
	Tid      string
	Rank     int
	League   int
	Promoted bool
}

type TournamentScore struct {
	Uid   string
	Score uint64
	Data  string
}

type TournamentScoreData struct {
	CountryCode string   `json:"country_code"`
	Duration    uint64   `json:"duration"`
	Heroes      []int    `json:"heroes"`
	KillCount   int      `json:"kill_count"`
	Levels      []int    `json:"levels"`
	Ranks       []int    `json:"ranks"`
	Skins       []string `json:"skins"`
	Username    string   `json:"username"`
	W17Bonus    bool     `json:"w17_bonus"`
	W17Stars    int      `json:"w17_stars"`
}

type TournamentData struct {
	Maps       []TournamentMapData      `json:"maps"`
	Overrides  []TournamentOverrideData `json:"overrides"`
	EndTime    int                      `json:"end_time"`
	NextTime   int                      `json:"next_time"`
	MinVersion string                   `json:"min_version"`
}

type TournamentMapData struct {
	ItemCount     int             `json:"item_count"`
	Scene         string          `json:"scene"`
	Seed          int             `json:"seed"`
	Promotion     int             `json:"promotion"`
	DisabledItems []string        `json:"disabled_items"`
	Rewards       [][]interface{} `json:"rewards"`
	OverrideIndex int             `json:"override_index"`
}

type TournamentOverrideData struct {
	MaxLife        int                          `json:"max_life"`
	EnemyCount     int                          `json:"enemy_count"`
	WaveReward     int                          `json:"wave_reward"`
	RewardCap      int                          `json:"reward_cap"`
	InitialGold    int                          `json:"initial_gold"`
	Hero           TournamentOverrideHeroData   `json:"hero"`
	TowerModifiers [][]int                      `json:"tower_modifiers"`
	Waves          []TournamentOverrideWaveData `json:"waves"`
}

type TournamentOverrideHeroData struct {
	Index     int     `json:"index"`
	Modifiers [][]int `json:"modifiers"`
}

type TournamentOverrideWaveData struct {
	Index      string                             `json:"index"`
	EnemyCount int                                `json:"enemy_count"`
	Flavors    [][]string                         `json:"flavors"`
	Modifiers  TournamentOverrideWaveModifierData `json:"modifiers"`
}

type TournamentOverrideWaveModifierData struct {
	Hp       int     `json:"hp"`
	Spd      float64 `json:"spd"`
	Atk      int     `json:"atk"`
	HpScale  float64 `json:"hp_scale"`
	SpdScale float64 `json:"spd_scale"`
	AtkScale float64 `json:"atk_scale"`
}

type SaveRequest struct {
	Id   string
	Seq  int
	Data string
}

type SaveData struct {
	// see extrajson.go
	Extra map[string]interface{} `json:"-"`

	Acd SaveDataAcd `json:"acd"`
	// Ads are the watched-ads status
	Ads   SaveDataAds   `json:"ads"`
	Btsdk SaveDataBtSdk `json:"btsdk"`
	// Cc is the player's country code
	Cc  string        `json:"cc"`
	Cs  []string      `json:"cs"`
	Cw  int           `json:"cw"`
	Cwm int           `json:"cwm"`
	Di  int           `json:"di"`
	Dlc []SaveDataDlc `json:"dlc"`
	Dq  SaveDataDq    `json:"dq"`
	// E is the player's energy
	E int `json:"e"`
	// El is the information about the infinite levels in each world
	El SaveDataEl `json:"el"`
	// Em is the information about the elixir mine
	Em []SaveDataEm `json:"em"`
	// Et is ??? timestamp
	Et  uint64
	Evt SaveDataEvt `json:"evt"`
	// Exp are the UI experiments to make you spend money on the game
	Exp SaveDataExp `json:"exp"`
	// F is some facebook info? It's empty for me.
	F    []interface{} `json:"f"`
	Fbc  bool          `json:"fbc"`
	Fbgt int           `json:"fbgt"`
	Fek  SaveDataFek   `json:"fek"`
	Fnd  []interface{} `json:"fnd"`
	// Fs are the Forge Slots
	Fs []SaveDataFs `json:"fs"`
	// Hs is the heroes data. It's an array of array of various values. Here's a quick rundown of the layout:
	// [
	//		Hero Index [int],
	//		??? always 3 or 4 for me [int],
	//		OSV value of the hero's level,
	//		OSV value of the current level XP progress,
	//		[
	//			OSV value of the first ability level,
	//			OSV value of the second ability level,
	//			OSV value of the third ability level,
	//			OSV value of the fourth ability level,
	//		],
	//		[
	//			the story progress (e.g. values: "1.1", "1.2", "2.1", "2.2", ...), when leveling up abilities [string]
	//		],
	//		the name of the active skin, or "" if no skin set [string],
	//		[
	//			the name of an unlocked skin, or "" if none [string],
	//		],
	//		OSV value of the awakening rank,
	//
	//
	Hs [][]interface{} `json:"hs"`
	// Iv is the items values
	Iv  SaveDataIv `json:"iv"`
	Jdv int
	Jt  int
	Jv  string
	// Live are the events values
	Live SaveDataLive `json:"live"`
	Lpt  int          `json:"lpt"`
	// Mails are the status of in-game mails
	Mails SaveDataMails `json:"mails"`
	Mem   SaveDataMem   `json:"mem"`
	Mh    int           `json:"mh"`
	Mq    SaveDataMq    `json:"mq"`
	Pa    float64       `json:"pa"`
	// Q is the quests data
	Q   SaveDataQuests `json:"q"`
	Rvw bool           `json:"rvw"`
	Rwd SaveDataRwd    `json:"rwd"`
	// Sal is the Player's nickname
	Sal string `json:"sal"`
	// Seq is the sequence number of the save. Each new save increases the seq.
	Seq int `json:"seq"`
	// Set is the game settings
	Set SaveDataSet `json:"set"`
	// Sid is the player's Google Play Games session ID
	Sid   string        `json:"sid"`
	Siege SaveDataSiege `json:"siege"`
	// Spl is the Session/social play name (on Android, this value is "PlayGames")
	Spl string `json:"spl"`
	// Sps is a string array of the selected items (me for meteor, r for revival pot, etc see items)
	Sps []string `json:"sps"`
	// Sr is the Shattered Realms data
	Sr SaveDataSr `json:"sr"`
	Ti int        `json:"ti"`
	// Tlf is the last tournament season played? Something like that?
	Tlf string      `json:"tlf"`
	Tou SaveDataTou `json:"tou"`
	// Tut is the tutorial progress
	Tut SaveDataTut `json:"tut"`
	// Tz seems to be the UTC offset (timezone)
	Tz float64 `json:"tz"`
	// Uid is the player's UUID
	Uid string `json:"uid"`
	Un  string `json:"un"`
	Unv string `json:"unv"`
	V   int    `json:"v"`
	// WS []SaveDataWs
}

type SaveDataTut struct {
	// Pl is the tutorials labels and whether or not the played has seen them
	Pl map[string]bool `json:"pl"`
	Ps map[string]bool `json:"ps"`
}

type SaveDataTou struct {
	Bs  int `json:"bs"`
	Lp  int `json:"lp"`
	Lps int `json:"lps"`
	Pc  int `json:"pc"`
	Pt  int `json:"pt"`
}

// SaveDataSr is the Shattered Realms data
type SaveDataSr struct {
	Fp SaveDataSrFp `json:"fp"`
	// Lc is an int array, one value for each SR level. Everything is 0 for me so I'm not sure what they mean.
	Lc []int `json:"lc"`
	// Ls is the max reached level progress (game shows up to 5/5, but it can actually go up to 6)
	Ls []int `json:"ls"`
}

type SaveDataSrFp struct {
	I int `json:"i"`
	// T is the timestamp of ???
	T uint64 `json:"t"`
}

type SaveDataSiege struct {
	Bc    SaveDataOsvVal `json:"bc"`
	Ch    int            `json:"ch"`
	Count SaveDataOsvVal `json:"count"`
	Lbt   SaveDataOsvVal `json:"lbt"`
	Lt    SaveDataOsvVal `json:"lt"`
	MdCl  SaveDataOsvVal `json:"md_cl"`
	Pc    SaveDataOsvVal `json:"pc"`
	Plays string         `json:"plays"`
}

type SaveDataSet struct {
	Cml  bool   `json:"cml"`
	Drag bool   `json:"drag"`
	Lang string `json:"lang"`
	Mus  bool   `json:"mus"`
	Noti bool   `json:"noti"`
	Q    int    `json:"q"`
	Sfx  bool   `json:"sfx"`
	Vbr  bool   `json:"vbr"`
}

type SaveDataRwd struct {
	D int `json:"d"`
	I int `json:"i"`
	T int `json:"t"`
}

type SaveDataQuests struct {
	Dlc []bool `json:"dlc"`
	// Ws are the progress of world levels
	Ws []SaveDataQuestsWs `json:"ws"`
}

type SaveDataQuestsWs struct {
	// C is a string with 1 and 0 indicating whether the level was completed in normal mode with 3 stars
	C string `json:"c"`
	// L is a string with 1 and 0 indicating whether the level was completed in legendary mode
	L string `json:"l"`
}

type SaveDataMq struct {
	Fb  bool    `json:"fb"`
	Fbv bool    `json:"fbv"`
	H   int     `json:"h"`
	Tn  int     `json:"tn"`
	Tnr []*bool `json:"tnr"`
}

type SaveDataMem struct {
	M14   SaveDataMemEntry `json:"m14"`
	M14v2 SaveDataMemEntry `json:"m14v2"`
}

type SaveDataMemEntry struct {
	A bool `json:"a"`
	D int  `json:"d"`
	I int  `json:"i"`
	P int  `json:"p"`
}

type SaveDataMails struct {
	// Deleted is an array of the timestamps(?) of each deleted mail
	Deleted []uint64 `json:"deleted"`
	// Read is an array of the timestamps(?) of each read mail
	Read []uint64 `json:"read"`
}

type SaveDataLive struct {
	Caldera    SaveDataLiveCaldera    `json:"caldera"`
	DivineStar SaveDataLiveDivineStar `json:"divine_star"`
	Kaguya     SaveDataLiveKaguya     `json:"kaguya"`
}

type SaveDataLiveCaldera struct {
	Boxes   SaveDataOsvVal `json:"boxes"`
	Bundles SaveDataOsvVal `json:"bundles"`
	Claimed SaveDataOsvVal `json:"claimed"`
	Lv1     SaveDataOsvVal `json:"lv1"`
	Lv2     SaveDataOsvVal `json:"lv2"`
	Lv3     SaveDataOsvVal `json:"lv3"`
}

type SaveDataLiveKaguya struct {
	Boxes   SaveDataOsvVal `json:"boxes"`
	Bundles SaveDataOsvVal `json:"bundles"`
	Claimed SaveDataOsvVal `json:"claimed"`
	Lv1     SaveDataOsvVal `json:"lv1"`
	Lv2     SaveDataOsvVal `json:"lv2"`
	Lv3     SaveDataOsvVal `json:"lv3"`
}

type SaveDataLiveDivineStar struct {
	Boxes   SaveDataOsvVal `json:"boxes"`
	Bundles SaveDataOsvVal `json:"bundles"`
	Lv1     SaveDataOsvVal `json:"lv1"`
	Lv2     SaveDataOsvVal `json:"lv2"`
	Lv3     SaveDataOsvVal `json:"lv3"`
}

// SaveDataIv is the items values
type SaveDataIv struct {
	// Am is the number of Armageddon
	Am SaveDataIvEntry `json:"am"`
	// B is the number of fire bombs
	B SaveDataIvEntry `json:"b"`
	// DivineStar is the number of winter-event divine stars
	DivineStar SaveDataIvEntry `json:"divine_star"`
	// Ek is the number of endless run keys
	Ek SaveDataIvEntry `json:"ek"`
	// EvtTkCaldera is the number of Caldera tokens from the August 2018 Caldera event (hello 7500 gem bundle
	// which actually only used 750 gems!)
	EvtTkCaldera SaveDataIvEntry `json:"evt_tk_caldera"`
	// EvtTkKaguya is the number of Azura tokens (why is she called Kaguya everywhere?)
	EvtTkKaguya SaveDataIvEntry `json:"evt_tk_kaguya"`
	// F is the number of freeze potions
	F SaveDataIvEntry `json:"f"`
	// FF is the number of level 2 free potions
	FF SaveDataIvEntry `json:"ff"`
	// G is the number of gems
	G SaveDataIvEntry `json:"g"`
	// Me is the number of meteors
	Meteors SaveDataIvEntry `json:"me"`
	// R is the number of revival/heal potions
	R SaveDataIvEntry `json:"r"`
	// Rt is the number of Raid Tickets
	Rt SaveDataIvEntry `json:"rt"`
	// S is the number of summon potions
	S SaveDataIvEntry `json:"s"`
	// Siegekey is the number of Realm Siege keys
	Siegekey SaveDataIvEntry `json:"siegekey"`
	// Siegemd is the number of Realm Siege medals
	Siegemd SaveDataIvEntry `json:"siegemd"`
	// Tbolton is the number of Bolton Awakening Tokens
	Tbolton SaveDataIvEntry `json:"tbolton"`
	// Tcaldera is the number of Caldera Awakening Tokens
	Tcaldera SaveDataIvEntry `json:"tcaldera"`
	// Tconnie is the number of Connie Awakening Tokens
	Tconnie SaveDataIvEntry `json:"tconnie"`
	// Tefrigid is the number of Efrigid Awakening Tokens
	Tefrigid SaveDataIvEntry `json:"tefrigid"`
	// Tfee is the number of Fee Awakening Tokens
	Tfee SaveDataIvEntry `json:"tfee"`
	// Thelios is the number of Helios Awakening Tokens
	Thelios SaveDataIvEntry `json:"thelios"`
	// Thogan is the number of Hogan Awakening Tokens
	Thogan SaveDataIvEntry `json:"thogan"`
	// Tkaguya is the number of Azura Awakening Tokens
	Tkaguya SaveDataIvEntry `json:"tkaguya"`
	// Tlancelot is the number of Lancelot Awakening Tokens
	Tlancelot SaveDataIvEntry `json:"tlancelot"`
	// Tleif is the number of Leif Awakening Tokens
	Tleif SaveDataIvEntry `json:"tleif"`
	// Tmabyn is the number of Mabyn Awakening Tokens
	Tmabyn SaveDataIvEntry `json:"tmabyn"`
	// Tmasamune is the number of Masamune Awakening Tokens
	Tmasamune SaveDataIvEntry `json:"tmasamune"`
	// Tnarlax is the number of Narlax Awakening Tokens
	Tnarlax SaveDataIvEntry `json:"tnarlax"`
	// Tobsidian is the number of Obsidian Awakening Tokens
	Tobsidian SaveDataIvEntry `json:"tobsidian"`
	// Tsethos is the number of Sethos Awakening Tokens
	Tsethos SaveDataIvEntry `json:"tsethos"`
	// Tsmoulder is the number of Smoulder Awakening Tokens
	Tsmoulder SaveDataIvEntry `json:"tsmoulder"`
	// Tyan is the number of Yan Awakening Tokens
	Tyan SaveDataIvEntry `json:"tyan"`
	// Wska is the number of ?? (I have 50)
	Wska SaveDataIvEntry `json:"wska"`
	// Wskb is the number of ?? (I have 120)
	Wskb SaveDataIvEntry `json:"wskb"`
	// Wskc is the number of ?? (I have 500)
	Wskc SaveDataIvEntry `json:"wskc"`
	// Wskd is the number of ?? (I have 1000)
	Wskd SaveDataIvEntry `json:"wskd"`
	// Wspa is the number of ?? (I have 500)
	Wspa SaveDataIvEntry `json:"wspa"`
	// Wspb is the number of ?? (I have 1500)
	Wspb SaveDataIvEntry `json:"wspb"`
	// Wstar is the number of ?? (I have 0)
	Wstar SaveDataIvEntry `json:"wstar"`
	// X is the amount of elixir
	X SaveDataIvEntry `json:"x"`
}

// SaveDataIvEntry is another awful attempt at obfuscating integer values. See GetValueFromIVEntry for details.
type SaveDataIvEntry struct {
	G SaveDataOsvVal `json:"g"`
	U SaveDataOsvVal `json:"u"`
}

// SaveDataFs is a forge slot
type SaveDataFs struct {
	// Ad is whether an ad has been watched to forward the time
	Ad int `json:"ad"`
	// K is the type of item brewing in the slot (am = armageddon)
	K string `json:"k"`
	// St is the timestamp (uint64) at which you started brewing the item, or "" if there was never any
	St interface{} `json:"st"`
	// U is whether the slot is currently used
	U bool `json:"u"`
}

// OSV-encoded ("obfuscation" lolol) of numerical values.
// Value example:
// {"o": -10124, "s": -1, "v": -10151}  =>  -10124 - (-10151) = 27, one of my hero level
type SaveDataOsvVal struct {
	O int `json:"o"`
	// S is either 1 or -1, it seems like a value by which O and V should be multiplied?
	S int `json:"s"`
	V int `json:"v"`
}

type SaveDataFek struct {
	C int `json:"c"`
	D int `json:"d"`
	I int `json:"i"`
	T int `json:"t"`
}

type SaveDataExp struct {
	BoltonTime      SaveDataExpEntry `json:"bolton_time"`
	BundleAutoPopup SaveDataExpEntry `json:"bundle_auto_popup"`
	IapNewStore     SaveDataExpEntry `json:"iap_new_store"`
	IapPrice        SaveDataExpEntry `json:"iap_price"`
	IapStarter      SaveDataExpEntry `json:"iap_starter"`
}

type SaveDataExpEntry struct {
	Value   int `json:"value"`
	Version int `json:"version"`
}

type SaveDataEvt struct {
	I int `json:"i"`
	L int `json:"l"`
	N int `json:"n"`
	T int `json:"t"`
	W int `json:"w"`
}

// SaveDataEm is the information about the elixir mine
type SaveDataEm struct {
	L int `json:"l"`
	// Lct is the last collection time
	Lct uint64 `json:"lct"`
	// Pl is the mine's level (0-indexed, so 7 = level 8)
	Pl int     `json:"pl"`
	S  float64 `json:"s"`
}

// SaveDataEl is the information about the endless levels in each world
type SaveDataEl struct {
	W1L21 SaveDataElEntry `json:"w1l21"`
	W2L21 SaveDataElEntry `json:"w2l21"`
	W3L21 SaveDataElEntry `json:"w3l21"`
	W4L21 SaveDataElEntry `json:"w4l21"`
	W5L21 SaveDataElEntry `json:"w5l21"`
	W6L21 SaveDataElEntry `json:"w6l21"`
}

type SaveDataElEntry struct {
	// Bs is the duration of the best run of the endless level
	Bs float64 `json:"bs"`
	// Lp is the timestamp at which the level was last played
	Lp uint64 `json:"lp"`
	Pc int    `json:"pc"`
	Q  string `json:"q"`
}

type SaveDataDq struct {
	D int `json:"d"`
	I int `json:"i"`
}

type SaveDataDlc struct {
	Lc []int `json:"lc"`
	Ls []int `json:"ls"`
}

type SaveDataAcd struct {
	Ls []int `json:"ls"`
	// Lt is the last timestamp
	Lt int `json:"lt"`
}

type SaveDataAds struct {
	Lb SaveDataAdsLb `json:"lb"`
	// Rt is the ads information about how many Raid Ticket ads you've watched
	Rt SaveDataAdsRaidTicket `json:"rt"`
	// Tn is the ads information about how many Tournament ads you've watched
	Tn SaveDataAdsTournament `json:"tn"`
}

type SaveDataAdsLb struct {
	On bool `json:"on"`
}

type SaveDataAdsRaidTicket struct {
	// I is the number of ads left to watch
	I int `json:"i"`
	// Lt is the timestamp at which the last ad was seen
	Lt uint64 `json:"lt"`
}

type SaveDataAdsTournament struct {
	// I is the number of ads left to watch
	I int `json:"i"`
	// Lt is the timestamp at which the last ad was seen
	Lt uint64 `json:"lt"`
	Pc int    `json:"pc"`
}

type SaveDataBtSdk struct {
	Fp   bool `json:"fp"`
	Fq   bool `json:"fq"`
	Lct  int  `json:"lct"`
	Lolt int  `json:"lolt"`
	Olt  int  `json:"olt"`
	Sp   bool `json:"sp"`
}

type StaticTnNewsRequest struct {
}

type StaticTnNewsResponse []struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	MinVersion string `json:"min_version"`
	Season     int    `json:"season"`
	Hero1      int    `json:"hero1"`
	Hero2      int    `json:"hero2"`
}

type TournamentScoreRequest struct {
	Data  string // TournamentScoreData
	Power int
	Realm string
	Score uint64
	Tid   string
	Uid   string
}
