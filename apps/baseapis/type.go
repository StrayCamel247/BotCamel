package baseapis

import (
// "fmt"
)

//	InfoMenifestBaseDB
type InfoMenifestBaseDB struct {
	ItemId      string `json:"itemid"`
	Description string `json:"description"`
	// Name        string `gorm:"SIZE:0;Name:name;index:,sort:desc,collate:utf8,type:btree"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Tag      string `json:"tag"`
	SeasonId string `json:"seasonid"`
}
type ItemIdDB struct {
	ItemId      string
	Description string
	Name        string
	Tag         string
}

// Manifest 返回数据结构体
type ManifestJsonFilePath struct {
	DestinyEnemyRaceDefinition                      string `json:"DestinyEnemyRaceDefinition"`
	DestinyNodeStepSummaryDefinition                string `json:"DestinyNodeStepSummaryDefinition"`
	DestinyArtDyeChannelDefinition                  string `json:"DestinyArtDyeChannelDefinition"`
	DestinyArtDyeReferenceDefinition                string `json:"DestinyArtDyeReferenceDefinition"`
	DestinyPlaceDefinition                          string `json:"DestinyPlaceDefinition"`
	DestinyActivityDefinition                       string `json:"DestinyActivityDefinition"`
	DestinyActivityTypeDefinition                   string `json:"DestinyActivityTypeDefinition"`
	DestinyClassDefinition                          string `json:"DestinyClassDefinition"`
	DestinyGenderDefinition                         string `json:"DestinyGenderDefinition"`
	DestinyInventoryBucketDefinition                string `json:"DestinyInventoryBucketDefinition"`
	DestinyRaceDefinition                           string `json:"DestinyRaceDefinition"`
	DestinyTalentGridDefinition                     string `json:"DestinyTalentGridDefinition"`
	DestinyUnlockDefinition                         string `json:"DestinyUnlockDefinition"`
	DestinyMaterialRequirementSetDefinition         string `json:"DestinyMaterialRequirementSetDefinition"`
	DestinySandboxPerkDefinition                    string `json:"DestinySandboxPerkDefinition"`
	DestinyStatGroupDefinition                      string `json:"DestinyStatGroupDefinition"`
	DestinyProgressionMappingDefinition             string `json:"DestinyProgressionMappingDefinition"`
	DestinyFactionDefinition                        string `json:"DestinyFactionDefinition"`
	DestinyVendorGroupDefinition                    string `json:"DestinyVendorGroupDefinition"`
	DestinyRewardSourceDefinition                   string `json:"DestinyRewardSourceDefinition"`
	DestinyUnlockValueDefinition                    string `json:"DestinyUnlockValueDefinition"`
	DestinyRewardMappingDefinition                  string `json:"DestinyRewardMappingDefinition"`
	DestinyRewardSheetDefinition                    string `json:"DestinyRewardSheetDefinition"`
	DestinyItemCategoryDefinition                   string `json:"DestinyItemCategoryDefinition"`
	DestinyDamageTypeDefinition                     string `json:"DestinyDamageTypeDefinition"`
	DestinyActivityModeDefinition                   string `json:"DestinyActivityModeDefinition"`
	DestinyMedalTierDefinition                      string `json:"DestinyMedalTierDefinition"`
	DestinyAchievementDefinition                    string `json:"DestinyAchievementDefinition"`
	DestinyActivityGraphDefinition                  string `json:"DestinyActivityGraphDefinition"`
	DestinyActivityInteractableDefinition           string `json:"DestinyActivityInteractableDefinition"`
	DestinyBondDefinition                           string `json:"DestinyBondDefinition"`
	DestinyCharacterCustomizationCategoryDefinition string `json:"DestinyCharacterCustomizationCategoryDefinition"`
	DestinyCharacterCustomizationOptionDefinition   string `json:"DestinyCharacterCustomizationOptionDefinition"`
	DestinyCollectibleDefinition                    string `json:"DestinyCollectibleDefinition"`
	DestinyDestinationDefinition                    string `json:"DestinyDestinationDefinition"`
	DestinyEntitlementOfferDefinition               string `json:"DestinyEntitlementOfferDefinition"`
	DestinyEquipmentSlotDefinition                  string `json:"DestinyEquipmentSlotDefinition"`
	DestinyStatDefinition                           string `json:"DestinyStatDefinition"`
	DestinyInventoryItemDefinition                  string `json:"DestinyInventoryItemDefinition"`
	DestinyInventoryItemLiteDefinition              string `json:"DestinyInventoryItemLiteDefinition"`
	DestinyItemTierTypeDefinition                   string `json:"DestinyItemTierTypeDefinition"`
	DestinyLocationDefinition                       string `json:"DestinyLocationDefinition"`
	DestinyLoreDefinition                           string `json:"DestinyLoreDefinition"`
	DestinyMetricDefinition                         string `json:"DestinyMetricDefinition"`
	DestinyObjectiveDefinition                      string `json:"DestinyObjectiveDefinition"`
	DestinyPlatformBucketMappingDefinition          string `json:"DestinyPlatformBucketMappingDefinition"`
	DestinyPlugSetDefinition                        string `json:"DestinyPlugSetDefinition"`
	DestinyPowerCapDefinition                       string `json:"DestinyPowerCapDefinition"`
	DestinyPresentationNodeDefinition               string `json:"DestinyPresentationNodeDefinition"`
	DestinyPresentationNodeBaseDefinition           string `json:"DestinyPresentationNodeBaseDefinition"`
	DestinyProgressionDefinition                    string `json:"DestinyProgressionDefinition"`
	DestinyProgressionLevelRequirementDefinition    string `json:"DestinyProgressionLevelRequirementDefinition"`
	DestinyRecordDefinition                         string `json:"DestinyRecordDefinition"`
	DestinyRewardAdjusterPointerDefinition          string `json:"DestinyRewardAdjusterPointerDefinition"`
	DestinyRewardAdjusterProgressionMapDefinition   string `json:"DestinyRewardAdjusterProgressionMapDefinition"`
	DestinyRewardItemListDefinition                 string `json:"DestinyRewardItemListDefinition"`
	DestinySackRewardItemListDefinition             string `json:"DestinySackRewardItemListDefinition"`
	DestinySandboxPatternDefinition                 string `json:"DestinySandboxPatternDefinition"`
	DestinySeasonDefinition                         string `json:"DestinySeasonDefinition"`
	DestinySeasonPassDefinition                     string `json:"DestinySeasonPassDefinition"`
	DestinySocketCategoryDefinition                 string `json:"DestinySocketCategoryDefinition"`
	DestinySocketTypeDefinition                     string `json:"DestinySocketTypeDefinition"`
	DestinyTraitDefinition                          string `json:"DestinyTraitDefinition"`
	DestinyTraitCategoryDefinition                  string `json:"DestinyTraitCategoryDefinition"`
	DestinyUnlockCountMappingDefinition             string `json:"DestinyUnlockCountMappingDefinition"`
	DestinyUnlockEventDefinition                    string `json:"DestinyUnlockEventDefinition"`
	DestinyUnlockExpressionMappingDefinition        string `json:"DestinyUnlockExpressionMappingDefinition"`
	DestinyVendorDefinition                         string `json:"DestinyVendorDefinition"`
	DestinyMilestoneDefinition                      string `json:"DestinyMilestoneDefinition"`
	DestinyActivityModifierDefinition               string `json:"DestinyActivityModifierDefinition"`
	DestinyReportReasonCategoryDefinition           string `json:"DestinyReportReasonCategoryDefinition"`
	DestinyArtifactDefinition                       string `json:"DestinyArtifactDefinition"`
	DestinyBreakerTypeDefinition                    string `json:"DestinyBreakerTypeDefinition"`
	DestinyChecklistDefinition                      string `json:"DestinyChecklistDefinition"`
	DestinyEnergyTypeDefinition                     string `json:"DestinyEnergyTypeDefinition"`
}
type ManifestLanguages struct {
	// 简体中文
	ZhChs ManifestJsonFilePath `json:"zh-chs"`
	// 繁体中文
	ZhCht ManifestJsonFilePath `json:"zh-cht"`
	// 英文
	En ManifestJsonFilePath `json:"en"`
}
type ManifestWorldComponentContent struct {
	JsonWorldComponentContentPaths ManifestLanguages `json:"jsonWorldComponentContentPaths"`
	NewVersion                     string            `json:"version"`
}
type ManifestResult struct {
	Response ManifestWorldComponentContent `json:"Response"`
}

type Properties struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
}
type InfoDisplay struct {
	// Properties map[string]string `json:"displayProperties"`
	Properties Properties `json:"displayProperties"`
	SeasonHash int64      `json:"seasonHash"`
}

// AccountStats 用户生涯记录查询数据结构体
type AccountStatsVlue struct {
	Value        float64 `json:"value"`
	DisplayValue string  `json:"displayValue"`
}
type AccountStatsInfo struct {
	StatId string           `json:"statId"`
	Basic  AccountStatsVlue `json:"basic"`
	Pga    AccountStatsVlue `json:"pga"`
}
type AccountStatsAllItems struct {
	ActivitiesEntered             AccountStatsInfo `json:"activitiesEntered"`
	ActivitiesWon                 AccountStatsInfo `json:"activitiesWon"`
	Assists                       AccountStatsInfo `json:"assists"`
	TotalDeathDistance            AccountStatsInfo `json:"totalDeathDistance"`
	AverageDeathDistance          AccountStatsInfo `json:"averageDeathDistance"`
	TotalKillDistance             AccountStatsInfo `json:"totalKillDistance"`
	Kills                         AccountStatsInfo `json:"kills"`
	AverageKillDistance           AccountStatsInfo `json:"averageKillDistance"`
	SecondsPlayed                 AccountStatsInfo `json:"secondsPlayed"`
	Deaths                        AccountStatsInfo `json:"deaths"`
	AverageLifespan               AccountStatsInfo `json:"averageLifespan"`
	Score                         AccountStatsInfo `json:"score"`
	AverageScorePerKill           AccountStatsInfo `json:"averageScorePerKill"`
	AverageScorePerLife           AccountStatsInfo `json:"averageScorePerLife"`
	BestSingleGameKills           AccountStatsInfo `json:"bestSingleGameKills"`
	BestSingleGameScore           AccountStatsInfo `json:"bestSingleGameScore"`
	OpponentsDefeated             AccountStatsInfo `json:"opponentsDefeated"`
	Efficiency                    AccountStatsInfo `json:"efficiency"`
	KillsDeathsRatio              AccountStatsInfo `json:"killsDeathsRatio"`
	KillsDeathsAssists            AccountStatsInfo `json:"killsDeathsAssists"`
	ObjectivesCompleted           AccountStatsInfo `json:"objectivesCompleted"`
	PrecisionKills                AccountStatsInfo `json:"precisionKills"`
	ResurrectionsPerformed        AccountStatsInfo `json:"resurrectionsPerformed"`
	ResurrectionsReceived         AccountStatsInfo `json:"resurrectionsReceived"`
	Suicides                      AccountStatsInfo `json:"suicides"`
	WeaponKillsAutoRifle          AccountStatsInfo `json:"weaponKillsAutoRifle"`
	WeaponKillsBeamRifle          AccountStatsInfo `json:"weaponKillsBeamRifle"`
	WeaponKillsBow                AccountStatsInfo `json:"weaponKillsBow"`
	WeaponKillsFusionRifle        AccountStatsInfo `json:"weaponKillsFusionRifle"`
	WeaponKillsHandCannon         AccountStatsInfo `json:"weaponKillsHandCannon"`
	WeaponKillsTraceRifle         AccountStatsInfo `json:"weaponKillsTraceRifle"`
	WeaponKillsMachineGun         AccountStatsInfo `json:"weaponKillsMachineGun"`
	WeaponKillsPulseRifle         AccountStatsInfo `json:"weaponKillsPulseRifle"`
	WeaponKillsRocketLauncher     AccountStatsInfo `json:"weaponKillsRocketLauncher"`
	WeaponKillsScoutRifle         AccountStatsInfo `json:"weaponKillsScoutRifle"`
	WeaponKillsShotgun            AccountStatsInfo `json:"weaponKillsShotgun"`
	WeaponKillsSniper             AccountStatsInfo `json:"weaponKillsSniper"`
	WeaponKillsSubmachinegun      AccountStatsInfo `json:"weaponKillsSubmachinegun"`
	WeaponKillsRelic              AccountStatsInfo `json:"weaponKillsRelic"`
	WeaponKillsSideArm            AccountStatsInfo `json:"weaponKillsSideArm"`
	WeaponKillsSword              AccountStatsInfo `json:"weaponKillsSword"`
	WeaponKillsAbility            AccountStatsInfo `json:"weaponKillsAbility"`
	WeaponKillsGrenade            AccountStatsInfo `json:"weaponKillsGrenade"`
	WeaponKillsGrenadeLauncher    AccountStatsInfo `json:"weaponKillsGrenadeLauncher"`
	WeaponKillsSuper              AccountStatsInfo `json:"weaponKillsSuper"`
	WeaponKillsMelee              AccountStatsInfo `json:"weaponKillsMelee"`
	WeaponBestType                AccountStatsInfo `json:"weaponBestType"`
	WinLossRatio                  AccountStatsInfo `json:"winLossRatio"`
	AllParticipantsCount          AccountStatsInfo `json:"allParticipantsCount"`
	AllParticipantsScore          AccountStatsInfo `json:"allParticipantsScore"`
	AllParticipantsTimePlayed     AccountStatsInfo `json:"allParticipantsTimePlayed"`
	LongestKillSpree              AccountStatsInfo `json:"longestKillSpree"`
	LongestSingleLife             AccountStatsInfo `json:"longestSingleLife"`
	MostPrecisionKills            AccountStatsInfo `json:"mostPrecisionKills"`
	OrbsDropped                   AccountStatsInfo `json:"orbsDropped"`
	OrbsGathered                  AccountStatsInfo `json:"orbsGathered"`
	RemainingTimeAfterQuitSeconds AccountStatsInfo `json:"remainingTimeAfterQuitSeconds"`
	TeamScore                     AccountStatsInfo `json:"teamScore"`
	TotalActivityDurationSeconds  AccountStatsInfo `json:"totalActivityDurationSeconds"`
	CombatRating                  AccountStatsInfo `json:"combatRating"`
	FastestCompletionMs           AccountStatsInfo `json:"fastestCompletionMs"`
	LongestKillDistance           AccountStatsInfo `json:"longestKillDistance"`
	HighestCharacterLevel         AccountStatsInfo `json:"highestCharacterLevel"`
	HighestLightLevel             AccountStatsInfo `json:"highestLightLevel"`
	FireTeamActivities            AccountStatsInfo `json:"fireTeamActivities"`
	ActivitiesCleared             AccountStatsInfo `json:"activitiesCleared"`
	HeroicPublicEventsCompleted   AccountStatsInfo `json:"heroicPublicEventsCompleted"`
	AdventuresCompleted           AccountStatsInfo `json:"adventuresCompleted"`
	PublicEventsCompleted         AccountStatsInfo `json:"publicEventsCompleted"`
}
type AccountStatsAllTime struct {
	AllTime AccountStatsAllItems `json:"allTime"`
}
type AccountStatsAllPvPPvE struct {
	AllPvP AccountStatsAllTime `json:"allPvP"`
	AllPvE AccountStatsAllTime `json:"allPvE"`
}
type AccountStatsResults struct {
	Results AccountStatsAllPvPPvE `json:"results"`
}
type AccountStatsAllCharacters struct {
	MergedAllCharacters AccountStatsResults `json:"mergedAllCharacters"`
}
type AccountStatsResult struct {
	Response AccountStatsAllCharacters `json:"Response"`
}

// Baseprofile 用户生涯记录查询数据结构体
type BaseprofileUserInfoEle struct {
	DisplayName string `json:"displayName"`
}
type BaseprofileUserInfo struct {
	UserInfo BaseprofileUserInfoEle `json:"userInfo"`
}
type BaseprofileData struct {
	Data BaseprofileUserInfo `json:"data"`
}
type BaseprofileProfile struct {
	Profile BaseprofileData `json:"profile"`
}
type BaseprofileResult struct {
	Response BaseprofileProfile `json:"Response"`
}

func init() {
	// PvpInfoArray = []Info{
	// 	Info{name: "熔炉", itemId: "1618576679"},
	// 	Info{name: "智谋", itemId: "1556658903"},
	// }

}
