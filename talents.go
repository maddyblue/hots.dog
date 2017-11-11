package main

type Hero struct {
	Name      string
	Slug      string
	Role      string
	MultiRole []string
}

var heroData = []Hero{
	{
		Name:      "Abathur",
		Slug:      "abathur",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Alarak",
		Slug:      "alarak",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Ana",
		Slug:      "ana",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Anub'arak",
		Slug:      "anubarak",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Artanis",
		Slug:      "artanis",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Arthas",
		Slug:      "arthas",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Auriel",
		Slug:      "auriel",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Azmodan",
		Slug:      "azmodan",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Brightwing",
		Slug:      "brightwing",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Cassia",
		Slug:      "cassia",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Chen",
		Slug:      "chen",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Cho",
		Slug:      "cho",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Chromie",
		Slug:      "chromie",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "D.Va",
		Slug:      "dva",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Dehaka",
		Slug:      "dehaka",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Diablo",
		Slug:      "diablo",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "E.T.C.",
		Slug:      "etc",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Falstad",
		Slug:      "falstad",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Gall",
		Slug:      "gall",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Garrosh",
		Slug:      "garrosh",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Gazlowe",
		Slug:      "gazlowe",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Genji",
		Slug:      "genji",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Greymane",
		Slug:      "greymane",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Gul'dan",
		Slug:      "guldan",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Illidan",
		Slug:      "illidan",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Jaina",
		Slug:      "jaina",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Johanna",
		Slug:      "johanna",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Junkrat",
		Slug:      "junkrat",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Kael'thas",
		Slug:      "kaelthas",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Kel'Thuzad",
		Slug:      "kelthuzad",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Kerrigan",
		Slug:      "kerrigan",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Kharazim",
		Slug:      "kharazim",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Leoric",
		Slug:      "leoric",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Li Li",
		Slug:      "lili",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Li-Ming",
		Slug:      "liming",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Lt. Morales",
		Slug:      "ltmorales",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Lunara",
		Slug:      "lunara",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Lúcio",
		Slug:      "lucio",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Malfurion",
		Slug:      "malfurion",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Malthael",
		Slug:      "malthael",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Medivh",
		Slug:      "medivh",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Muradin",
		Slug:      "muradin",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Murky",
		Slug:      "murky",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Nazeebo",
		Slug:      "nazeebo",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Nova",
		Slug:      "nova",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Probius",
		Slug:      "probius",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Ragnaros",
		Slug:      "ragnaros",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Raynor",
		Slug:      "raynor",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Rehgar",
		Slug:      "rehgar",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Rexxar",
		Slug:      "rexxar",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Samuro",
		Slug:      "samuro",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Sgt. Hammer",
		Slug:      "sgthammer",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Sonya",
		Slug:      "sonya",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Stitches",
		Slug:      "stitches",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Stukov",
		Slug:      "stukov",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Sylvanas",
		Slug:      "sylvanas",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Tassadar",
		Slug:      "tassadar",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "The Butcher",
		Slug:      "thebutcher",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "The Lost Vikings",
		Slug:      "thelostvikings",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Thrall",
		Slug:      "thrall",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Tracer",
		Slug:      "tracer",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Tychus",
		Slug:      "tychus",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Tyrael",
		Slug:      "tyrael",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Tyrande",
		Slug:      "tyrande",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Uther",
		Slug:      "uther",
		Role:      "Support",
		MultiRole: []string{"Support"},
	},
	{
		Name:      "Valeera",
		Slug:      "valeera",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Valla",
		Slug:      "valla",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Varian",
		Slug:      "varian",
		Role:      "Multiclass",
		MultiRole: []string{"Warrior", "Damage"},
	},
	{
		Name:      "Xul",
		Slug:      "xul",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Zagara",
		Slug:      "zagara",
		Role:      "Specialist",
		MultiRole: []string{"Specialist"},
	},
	{
		Name:      "Zarya",
		Slug:      "zarya",
		Role:      "Warrior",
		MultiRole: []string{"Warrior"},
	},
	{
		Name:      "Zeratul",
		Slug:      "zeratul",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
	{
		Name:      "Zul'jin",
		Slug:      "zuljin",
		Role:      "Damage",
		MultiRole: []string{"Damage"},
	},
}

type talentText struct {
	Name string
	Text string
}

var talentData = map[string]talentText{
	"AbathurCombatStyleAssaultStrain": {
		Name: "Assault Strain",
		Text: "Locust Basic Attacks cleave for 50% damage, and explode on death for 102 damage.",
	},
	"AbathurCombatStyleBombardStrain": {
		Name: "Bombard Strain",
		Text: "Locust's Basic Attacks become a long-range siege attack that deal 70% more damage.",
	},
	"AbathurCombatStyleLocustBrood": {
		Name: "Locust Brood",
		Text: "Activate to spawn 3 Locusts at a nearby location.",
	},
	"AbathurCombatStyleSurvivalInstincts": {
		Name: "Survival Instincts",
		Text: "Increases Locust's Health by 50% and duration by 40%.",
	},
	"AbathurHeroicAbilityEvolveMonstrosity": {
		Name: "Evolve Monstrosity",
		Text: "Turn an allied Minion or Locust into a Monstrosity. When enemy Minions near the Monstrosity die, it gains 5% Health and 5% Basic Attack damage, stacking up to 40 times.  The Monstrosity takes 50% less damage from non-Heroic enemies.\nUsing Symbiote on the Monstrosity allows Abathur to control it, in addition to Symbiote's normal benefits.  This Ability can be reactivated to automatically cast Symbiote on his Monstrosity.",
	},
	"AbathurHeroicAbilityUltimateEvolution": {
		Name: "Ultimate Evolution",
		Text: "Clone target allied Hero and control it for 20 seconds. Abathur has perfected the clone, granting it 20% Spell Power, 20% bonus Attack Damage, and 10% bonus Movement Speed. Cannot use their Heroic Ability.",
	},
	"AbathurMasteryAdrenalineBoost": {
		Name: "Adrenaline Boost",
		Text: "Symbiote's Carapace increases the Movement Speed of the target by 40% for 3.5 seconds.",
	},
	"AbathurMasteryBallistospores": {
		Name: "Ballistospores",
		Text: "Increases Toxic Nest's range to global and increases duration by 25%.",
	},
	"AbathurMasteryEnvenomedNestsToxicNest": {
		Name: "Envenomed Nest",
		Text: "Toxic Nests deal 75% more damage over 3 seconds.",
	},
	"AbathurMasteryEnvenomedSpikes": {
		Name: "Envenomed Spikes",
		Text: "Abathur's Symbiote's Spike Burst also slows enemy Movement Speed by 40% for 2 seconds.",
	},
	"AbathurMasteryEvolutionComplete": {
		Name: "Evolution Complete",
		Text: "Monstrosity gains the ability to Deep Tunnel to any visible location once every 25 seconds.",
	},
	"AbathurMasteryEvolutionMasterUltimateEvolution": {
		Name: "Evolution Master",
		Text: "Evolution Master \nLower the cooldown of Ultimate Evolution to 90 seconds, and increase the duration to 60 seconds.",
	},
	"AbathurMasteryLocustMaster": {
		Name: "Locust Nest",
		Text: "Activate to create a nest that periodically spawns Locusts. Only one Locust Nest can be active at a time.",
	},
	"AbathurMasteryNeedlespine": {
		Name: "Needlespine",
		Text: "Increases the damage and range of Symbiote's Stab by 20%.",
	},
	"AbathurMasteryPressurizedGlands": {
		Name: "Pressurized Glands",
		Text: "Increases the range of Symbiote's Spike Burst by 25% and decreases the cooldown by 1 second.",
	},
	"AbathurMasteryProlificDispersal": {
		Name: "Prolific Dispersal",
		Text: "Reduces the cooldown of Toxic Nest by 2 seconds and adds 2 additional charges.",
	},
	"AbathurMasteryRegenerativeMicrobes": {
		Name: "Regenerative Microbes",
		Text: "Symbiote's Carapace heals the target for 68 Health per second over 4 seconds.",
	},
	"AbathurMasterySpatialEfficiency": {
		Name: "Spatial Efficiency",
		Text: "Symbiote's Stab gains 1 additional charge and its Cooldown is reduced by .5 seconds.",
	},
	"AbathurMasteryVileNestsToxicNest": {
		Name: "Vile Nest",
		Text: "Toxic Nests slow enemy Movement Speed by 40% for 2.5 seconds.",
	},
	"AbathurSymbioteAdrenalOverload": {
		Name: "Adrenal Overload",
		Text: "Heroic Symbiote hosts gain 25% increased Attack Speed.",
	},
	"AbathurSymbioteCarapaceNetworkedCarapace": {
		Name: "Networked Carapace",
		Text: "Using Symbiote's Carapace also applies an untalented Carapace Shield to all nearby allied Heroes, Minions, and Mercenaries.",
	},
	"AbathurSymbioteCarapaceSustainedCarapace": {
		Name: "Sustained Carapace",
		Text: "Increases the Shield amount of Symbiote's Carapace by 40% and allows it to persist after Symbiote ends.",
	},
	"AbathurSymbioteHivemind": {
		Name: "Hivemind",
		Text: "Symbiote creates an additional Symbiote on a nearby allied Hero.  This Symbiote mimics the commands of the first, but does half damage and shielding.",
	},
	"AbathurSymbioteSpikeBurstSomaTransference": {
		Name: "Soma Transference",
		Text: "Symbiote's Spike Burst heals the host for 64 Health per enemy Hero hit.",
	},
	"AbathurUltimateEvolutionEvolutionaryLink": {
		Name: "Evolutionary Link",
		Text: "As long as the Ultimate Evolution is alive, the original target of the clone gains a Shield equal to 25% of their maximum Health.  Refreshes every 5 seconds.",
	},
	"AbathurVolatileMutation": {
		Name: "Volatile Mutation",
		Text: "Ultimate Evolution clones and Monstrosities deal 137 damage to nearby enemies every 3 seconds and when they die.",
	},
	"AlarakAppliedForce": {
		Name: "Applied Force",
		Text: "Reduce Sadism by 10%, but Telekinesis pushes 20% further and has 20% increased range.",
	},
	"AlarakBladeoftheHighlord": {
		Name: "Blade of the Highlord",
		Text: "Basic Attacks against Heroes increase Sadism by 7% for 4 seconds, stacking up to 35%. Basic Attacks against Heroes refresh this duration.",
	},
	"AlarakChaosReigns": {
		Name: "Chaos Reigns",
		Text: "Quest: Hit Heroes with Discord Strike.\nReward: After hitting 15 Heroes, increase Discord Strike's damage by 50.\nReward: After hitting 2 Heroes with a single Discord Strike, increase its damage by 50.\nReward: After hitting 3 Heroes with a single Discord Strike, increase its damage by 80 and instantly gain all other Rewards.",
	},
	"AlarakCounterStrikeItem": {
		Name: "Counter-Strike",
		Text: "Alarak targets an area and channels for 1 second, becoming Protected and Unstoppable. After, if he took damage from an enemy Hero, he sends a shockwave that deals 275 damage.\nThis ability will take over Alarak's Trait button.",
	},
	"AlarakDeadlyChargeItem": {
		Name: "Deadly Charge",
		Text: "After channeling, Alarak charges forward dealing 200 damage to all enemies in his path. Distance is increased based on the amount of time channeled, up to 1.6 seconds.\nIssuing a Move order while this is channeling will cancel it at no cost. Taking damage will interrupt the channeling.\nThis ability will take over Alarak's Trait button.",
	},
	"AlarakDissonance": {
		Name: "Dissonance",
		Text: "Increase Discord Strike's Silence duration by 0.75 seconds.",
	},
	"AlarakExtendedLightning": {
		Name: "Extended Lightning",
		Text: "Reduce Sadism by 10%.\nQuest: Hit Heroes with the center of Lightning Surge.\nReward: After hitting 5 Heroes, increase Lightning Surge's range by 20%.\nReward: After hitting 15 Heroes, Lightning Surge's center also Slows enemies by 40% for 2 seconds.\nReward: After hitting 3 Heroes with the center of a single cast, increase Sadism by 10% and instantly gain all other Rewards.",
	},
	"AlarakHastyBargain": {
		Name: "Hasty Bargain",
		Text: "Activate to permanently reduce Alarak's Sadism by 4% and reset the cooldowns of his Basic Abilities.",
	},
	"AlarakHeroicAbilityCounterStrike": {
		Name: "Counter-Strike",
		Text: "Alarak targets an area and channels for 1 second, becoming Protected and Unstoppable. After, if he took damage from an enemy Hero, he sends a shockwave that deals 275 damage.",
	},
	"AlarakHeroicAbilityDeadlyCharge": {
		Name: "Deadly Charge",
		Text: "After channeling, Alarak charges forward dealing 200 damage to all enemies in his path. Distance is increased based on the amount of time channeled, up to 1.6 seconds.\nIssuing a Move order while this is channeling will cancel it at no cost. Taking damage will interrupt the channeling.",
	},
	"AlarakHinderedMotion": {
		Name: "Hindered Motion",
		Text: "Telekinesis Slows enemies by 25% for 2 seconds.",
	},
	"AlarakLastLaugh": {
		Name: "Last Laugh",
		Text: "Activate teleport to the chosen location and remove all Roots, Slows, and damage over time effects. If Alarak fails to hit enemy Heroes 3 times with Basic Abilities within 4 seconds of using Last Laugh, his Health is reduced to 1.",
	},
	"AlarakLethalOnslaughtDiscordStrike": {
		Name: "Lethal Onslaught",
		Text: "Hitting an enemy Hero with Discord Strike grants 50% of Sadism's damage bonus to Alarak's Basic Attacks for 4 seconds. Basic Attacks against Heroes refresh this duration.",
	},
	"AlarakLightningBarrage": {
		Name: "Lightning Barrage",
		Text: "Hitting an enemy Hero with the center of Lightning Surge allows it to be cast again for free on a different primary target within the next 2 seconds. This free cast cannot benefit from Lightning Barrage.",
	},
	"AlarakMockingStrikes": {
		Name: "Mocking Strikes",
		Text: "Basic Attacks against Stunned, Silenced, Slowed, or Rooted Heroes reduce the cooldown of Alarak's Basic Abilities by 2.5 seconds.",
	},
	"AlarakNegativelyCharged": {
		Name: "Negatively Charged",
		Text: "Increase the Lightning Surge damage bonus to enemies between Alarak and his victim to 150%.\nRepeatable Quest: Each Hero hit by the center of Lightning Surge permanently increases the damage bonus by 5%.",
	},
	"AlarakPureMalice": {
		Name: "Pure Malice",
		Text: "Increase Sadism by 10% any time an allied Hero dies, up to 40%.\nSadism gained from Pure Malice is lost on death.",
	},
	"AlarakRiteofRakShir": {
		Name: "Rite of Rak'Shir",
		Text: "Activate to mark an enemy Hero for 300 seconds.\nHitting the marked Hero with Discord Strike increases Sadism by 3%. Killing the marked Hero increases Sadism by 5% and lowers the cooldown of Rite of Rak'Shir to 10 seconds.\nSadism gained from Rite of Rak'Shir is lost on death.",
	},
	"AlarakRuthlessMomentum": {
		Name: "Ruthless Momentum",
		Text: "Basic Abilities recharge 20% faster while above 80% Health.",
	},
	"AlarakShowofForce": {
		Name: "Show of Force",
		Text: "Heroes hit with 3 of Alarak's Abilities within 2 seconds take an additional 77 damage.",
	},
	"AlarakSustainingPower": {
		Name: "Sustaining Power",
		Text: "Increase the healing received from damaging Heroes with Lightning Surge by 40%.",
	},
	"AmazonBallLightning": {
		Name: "Ball Lightning",
		Text: "Throw a ball of lightning at an enemy Hero that bounces up to 6 times between nearby enemy Heroes and Cassia, dealing 180 damage to enemies hit.",
	},
	"AmazonChargedStrikes": {
		Name: "Charged Strikes",
		Text: "Activate to increase Basic Attack damage by 20% and cause Basic Attacks against non-Structures to bounce to nearby enemy Heroes for 8 seconds.",
	},
	"AmazonGroundingBolt": {
		Name: "Grounding Bolt",
		Text: "Lightning Fury Slows enemies hit by 20% for 1.5 seconds.",
	},
	"AmazonImpale": {
		Name: "Impale",
		Text: "Enemies below 50% Health take 50% increased damage from Fend.",
	},
	"AmazonImprisoningLight": {
		Name: "Imprisoning Light",
		Text: "Upon impaling a Hero, Cassia's Valkyrie creates a wave of light that deals 200 damage to nearby enemies and Roots them for 3 seconds.",
	},
	"AmazonInfiniteLightning": {
		Name: "Infinite Lightning",
		Text: "Ball Lightning can now bounce indefinitely, and its cooldown is reduced by 3 seconds every time an enemy Hero is hit.",
	},
	"AmazonInnerLight": {
		Name: "Inner Light",
		Text: "Whenever Cassia is Stunned or Rooted, Blinding Light is cast at her location.\nThis effect has a 5.9 second cooldown.\nPassive: Blinding Light's radius is increased by 25%.",
	},
	"AmazonLungingStrike": {
		Name: "Lunging Strike",
		Text: "Increase Fend's range and area by 20%, and its duration by 0.5 seconds.",
	},
	"AmazonMartialLaw": {
		Name: "Martial Law",
		Text: "Basic Attacks against Stunned, Rooted, or Slowed enemy Heroes deal bonus damage equal to 3% of the Hero's maximum Health.",
	},
	"AmazonPenetrate": {
		Name: "Penetrate",
		Text: "Enemies take 3% increased damage from Fend each consecutive hit.",
	},
	"AmazonPierce": {
		Name: "Pierce",
		Text: "The primary missile of Lightning Fury now pierces, but splits only when hitting Heroes.",
	},
	"AmazonPlateoftheWhale": {
		Name: "Plate of the Whale",
		Text: "Regenerate 5 Health per second while Avoidance is active.\nQuest: For every 5 Basic Attacks absorbed with Avoidance, it regenerates an additional 1 Health per second.\nReward: After absorbing 75 Basic Attacks, increase Cassia's maximum Health by 10%.",
	},
	"AmazonRingoftheLeech": {
		Name: "Ring of the Leech",
		Text: "Cassia heals for 25% of the damage she deals to Blinded enemies.",
	},
	"AmazonSeraphsHymn": {
		Name: "Seraph's Hymn",
		Text: "Reduce Blinding Light's Mana cost from 70 to 40. Basic Attacks against Blinded enemies reduce the cooldown of Blinding Light by 2 seconds.",
	},
	"AmazonSurgeOfLight": {
		Name: "Surge of Light",
		Text: "After taking 500 damage with Avoidance active, Cassia can activate Avoidance to deal 225 damage to enemies around her.",
	},
	"AmazonThundergodsVigor": {
		Name: "Thundergod's Vigor",
		Text: "Enemy Heroes hit by the primary missile of Lightning Fury reduce its cooldown by 1 second.",
	},
	"AmazonThunderstroke": {
		Name: "Thunderstroke",
		Text: "Quest: Hit Heroes with Lightning Fury.\nReward: After hitting 20 Heroes, increase Lightning Fury's damage by 100.\nReward: After hitting 40 Heroes, Lightning Fury gains a 2nd charge and its Mana cost is reduced from 30 to 15.",
	},
	"AmazonTitansRevenge": {
		Name: "Titan's Revenge",
		Text: "Cassia's Basic Attacks now ignore Armor, and her Basic Attack range is increased by 2.",
	},
	"AmazonTrueSight": {
		Name: "True Sight",
		Text: "Quest: Damage Blinded enemy Heroes 60 times with Abilities or Basic Attacks.\nReward: Increase Blinding Light's Blind duration by 1 second.\nPassive: Increase Blinding Light's passive damage bonus by 20%.",
	},
	"AmazonValkyrie": {
		Name: "Valkyrie",
		Text: "Summon a Valkyrie that rushes to Cassia after 0.8 seconds, pulling the first enemy Hero hit, dealing 225 damage and Stunning them for 0.5 seconds at the end of her path. The Valkyrie knocks back all other enemy Heroes in her way.",
	},
	"AmazonWarTraveler": {
		Name: "War Traveler",
		Text: "Cassia gains 4% Movement Speed every 0.5 seconds while Avoidance is active, up to 20%.  This bonus is reset when Cassia stops moving or uses an Ability.",
	},
	"AnaAimDownSights": {
		Name: "Aim Down Sights",
		Text: "Shrike can be activated to increase vision radius by 100% and Basic Attack range by 2 but reduce Movement Speed by 20%. Lasts until canceled.",
	},
	"AnaAimDownSightsCustomOptics": {
		Name: "Custom Optics",
		Text: "Increase the Aim Down Sights Basic Attack range bonus from 2 to 4.",
	},
	"AnaBioticGrenadeAirStrike": {
		Name: "Air Strike",
		Text: "Activate to use Biotic Grenade with a 200% increased range, but Grenades thrown this way take 4 seconds to land.\nPassive: Reduce Biotic Grenade cooldown by 4 seconds.",
	},
	"AnaBioticGrenadeContactHealing": {
		Name: "Contact Healing",
		Text: "Biotic Grenade heals for 30% more per allied and enemy Hero hit.",
	},
	"AnaBioticGrenadeGrenadeCalibration": {
		Name: "Grenade Calibration",
		Text: "Quest: Hit enemy Heroes with Biotic Grenade.\nReward: After hitting 10 Heroes, Biotic Grenade does 75% more damage per allied and enemy Hero hit.\nReward: After hitting 20 Heroes, Biotic Grenade's duration on allies is increased from 4 seconds to 12 seconds, and its healing radius is increased by 100%.",
	},
	"AnaDebilitatingDart": {
		Name: "Debilitating Dart",
		Text: "Activate to fire a dart which reduces the damage dealt by the first enemy Hero it hits by 50% for 4 seconds.",
	},
	"AnaDetachableBoxMagazine": {
		Name: "Detachable Box Magazine",
		Text: "Quest: Stack 5 Doses on an enemy Hero.\nReward: After reaching maximum Dosage 5 times, Doses deal 50% increase damage to enemies with 5 Doses.\nReward: After reaching maximum Dosage 15 times, unlock the Active Reload Ability, which can be activated to instantly gain 3 charges of Healing Dart.",
	},
	"AnaDynamicShooting": {
		Name: "Dynamic Shooting",
		Text: "Basic Attacks increase Attack Speed by 10% for 4 seconds, up to 100%.",
	},
	"AnaEyeOfHorusBallisticAdvantage": {
		Name: "Ballistic Advantage",
		Text: "Eye of Horus rounds explodes upon impacting a Hero, healing nearby allies for 300 and dealing 150 damage to nearby enemy Heroes.",
	},
	"AnaHealingDartConcentratedDoses": {
		Name: "Concentrated Doses",
		Text: "Increase Healing Dart's healing by 10% for each Dose active on enemy Heroes.",
	},
	"AnaHealingDartPurifyingDarts": {
		Name: "Purifying Darts",
		Text: "Healing Dart removes Roots and Slows from the target, and heals for 20% more when doing so.",
	},
	"AnaHealingDartSharpshooter": {
		Name: "Sharpshooter",
		Text: "Healing Dart's healing is increased by 5% if it heals a Hero, up to 50%. This bonus is reset if Healing Dart fails to hit a Hero.",
	},
	"AnaHealingDartSmellingSalts": {
		Name: "Smelling Salts",
		Text: "Healing Dart removes Stuns from the target, and grants them 50 Armor for 2 seconds when doing so.",
	},
	"AnaHealingDartSpeedSerum": {
		Name: "Speed Serum",
		Text: "Healing Dart grants 25% Movement Speed to affected Heroes for 1.5 seconds.",
	},
	"AnaHeroicAbilityEyeOfHorus": {
		Name: "Eye of Horus",
		Text: "Assume a sniping position, gaining the ability to fire up to 8 specialized rounds with unlimited range. Rounds hit the first allied or enemy Hero or enemy Structure in their path. Allies are healed for 300 and enemies are damaged for 175. Deals 50% less damage to Structures.\nAna is unable to move while Eye of Horus is active.",
	},
	"AnaHeroicAbilityNanaBoost": {
		Name: "Nano Boost",
		Text: "Instantly boost an allied Hero, restoring 200 Mana. For the next 8 seconds, they gain 30% Spell Power and their Basic Ability cooldowns recharge 150% faster.\nCannot be used on Ana.",
	},
	"AnaMindNumbingAgent": {
		Name: "Mind-Numbing Agent",
		Text: "Every Dose a Hero has reduces their Spell Power by 15%.",
	},
	"AnaNanoInfusion": {
		Name: "Nano Infusion",
		Text: "Allies affected by Nano Boost heal for 50% of Spell Damage dealt.",
	},
	"AnaOverdose": {
		Name: "Overdose",
		Text: "Hitting enemy Heroes with Sleep Dart applies 3 Doses.",
	},
	"AnaPiercingDarts": {
		Name: "Piercing Darts",
		Text: "Quest: Hit Heroes with Sleep Dart.\nReward: After hitting 10 Heroes, Sleep Dart now hits 2 Heroes and its range is increased by 50%.\nReward: After hitting 20 Heroes, Healing Dart now hits 2 Heroes and its range is increased by 50%.",
	},
	"AnaSleepingDartTemporaryBlindness": {
		Name: "Temporary Blindness",
		Text: "After Sleep Dart's effects wear off, targets are Blinded for 2 seconds.",
	},
	"AnaSomnolentDoses": {
		Name: "Somnolent Doses",
		Text: "Upon stacking 5 Doses, Heroes are put to Sleep for 2 seconds.",
	},
	"AnubarakAcidDrenchedMandibles": {
		Name: "Acid Drenched Mandibles",
		Text: "Attacking a Hero that is slowed, rooted, or stunned increases Anub'arak's Basic Attack damage by 70% for 3 seconds.",
	},
	"AnubarakCocoonCryptweave": {
		Name: "Cryptweave",
		Text: "Staying near the Cocoon allows Anub'arak to extend the duration by up to 4 seconds per Cocoon.",
	},
	"AnubarakCombatStyleChitinousPlating": {
		Name: "Chitinous Plating",
		Text: "While Harden Carapace is active, taking damage from enemy Abilities reduces its cooldown by 1 second, up to 3 seconds.",
	},
	"AnubarakCombatStyleLegionOfBeetles": {
		Name: "Legion of Beetles",
		Text: "Automatically spawn a Beetle every 8 seconds.  Can be toggled on and off.",
	},
	"AnubarakDebilitation": {
		Name: "Debilitation",
		Text: "Enemy Heroes hit by Burrow Charge have their Spell Power reduced by 50% for 4 seconds.",
	},
	"AnubarakHeroicAbilityCarrionSwarm": {
		Name: "Locust Swarm",
		Text: "Deal 62 damage per second to nearby enemies. Each enemy damaged restores 21 Health. Lasts 6 seconds.",
	},
	"AnubarakHeroicAbilityCocoon": {
		Name: "Cocoon",
		Text: "Wraps target enemy Hero in a cocoon, rendering them unable to act or be targeted for 8 seconds. Allies of the Hero can attack the cocoon to break it and free them early.",
	},
	"AnubarakMasteryBedOfBarbs": {
		Name: "Bed of Barbs",
		Text: "Create a bed of spikes along Impale's path that slows enemies by 25% and deals 21 damage per second.  Spikes persist for 3.5 seconds.",
	},
	"AnubarakMasteryEpicenterBurrowCharge": {
		Name: "Epicenter",
		Text: "Increases Burrow Charge impact area by 60% and lowers the cooldown by 1.2 seconds for each Hero hit.",
	},
	"AnubarakMasteryHiveMasterCarrionSwarm": {
		Name: "Hive Master",
		Text: "Gain a permanent Vampire Locust that attacks a nearby enemy every 3 seconds. The Vampire Locust deals 153 damage and returns to heal Anub'arak for 37 health.",
	},
	"AnubarakMasteryLeechingScarabs": {
		Name: "Leeching Scarabs",
		Text: "Beetles gain 20% bonus attack damage and they heal Anub'arak for 50% of their damage with each attack if he is nearby.",
	},
	"AnubarakMasteryPersistentCarapaceHardenCarapace": {
		Name: "Persistent Carapace",
		Text: "Increases Harden Carapace's Shield duration by 3 seconds.",
	},
	"AnubarakMasteryShedExoskeletonHardenCarapace": {
		Name: "Shed Exoskeleton",
		Text: "Shed Exoskeleton\nGain 30% increased Move Speed for 3 seconds.",
	},
	"AnubarakMasteryUnderkingBurrowCharge": {
		Name: "Underking",
		Text: "Increases Burrow Charge range by 20% and damage by 100%.",
	},
	"AnubarakMasteryUrticatingSpines": {
		Name: "Urticating Spines",
		Text: "Casting Harden Carapace will also deal 80 damage to nearby enemies. Deals double damage against Heroes.",
	},
	"AnubarakNerubianArmor": {
		Name: "Nerubian Armor",
		Text: "Every 12 seconds, gain 30 Spell Armor against the next enemy Ability and subsequent Abilities for 1.5 seconds, reducing the damage taken by 30%.",
	},
	"AnubarakResilientScarabs": {
		Name: "Resilient Scarabs",
		Text: "Beetles have 50 Spell Armor.",
	},
	"AnubarakScarabHostBeetleJuiced": {
		Name: "Beetle, Juiced",
		Text: "Every 3rd Basic Attack against enemy Heroes spawns a Beetle.",
	},
	"AnubarakSubterraneanShield": {
		Name: "Subterranean Shield",
		Text: "Burrow Charge also grants a 345 point Shield for 5 seconds.",
	},
	"ArtanisBladeDashPsionicSynergy": {
		Name: "Psionic Synergy",
		Text: "Increases Blade Dash's cooldown reduction of Shield Overload by hitting Heroes by 2 seconds.",
	},
	"ArtanisBladeDashSolariteReaper": {
		Name: "Solarite Reaper",
		Text: "Increases the damage of the first dash of Blade Dash by 150%.",
	},
	"ArtanisBladeDashTemplarsZeal": {
		Name: "Templar's Zeal",
		Text: "Blade Dash cooldown recharges 150% faster while below 50% Health.",
	},
	"ArtanisHeroicAbilitySpearofAdunPurifierBeam": {
		Name: "Purifier Beam",
		Text: "Target an enemy Hero with an orbital beam from the Spear of Adun, dealing 184 damage per second for 8 seconds. The beam will chase the target as they move.  Unlimited range.",
	},
	"ArtanisHeroicAbilitySpearofAdunSuppressionPulse": {
		Name: "Suppression Pulse",
		Text: "Fire a large area pulse from the Spear of Adun, dealing 114 damage and Blinding enemies for 4 seconds. Unlimited range.",
	},
	"ArtanisPhasePrismChronoSurge": {
		Name: "Chrono Surge",
		Text: "Hitting an enemy Hero with Phase Prism grants 50% bonus Attack Speed for 5 seconds.",
	},
	"ArtanisPhasePrismGravitonVortex": {
		Name: "Graviton Vortex",
		Text: "Phase Prism pulls and damages an additional enemy Hero near the first, and its cooldown is reduced by 5 seconds.",
	},
	"ArtanisPhasePrismWarpSickness": {
		Name: "Warp Sickness",
		Text: "Phase Prism also slows the enemy's Movement Speed by 35% for 4 seconds.",
	},
	"ArtanisPlasmaBurn": {
		Name: "Plasma Burn",
		Text: "While Shield Overload is active, deal 55 damage per second to nearby enemies.",
	},
	"ArtanisShieldOverloadForceofWill": {
		Name: "Force of Will",
		Text: "Increases Shield Overload's cooldown reduction from Basic Attacks to 6 seconds.",
	},
	"ArtanisShieldOverloadPhaseBulwark": {
		Name: "Phase Bulwark",
		Text: "When Shield Overload activates, gain 40 Spell Armor for 4 seconds, reducing Ability Damage taken by 40%.",
	},
	"ArtanisShieldOverloadShieldBattery": {
		Name: "Shield Battery",
		Text: "Shield Overload's cooldown recharges 50% faster while its shield is active, and the shield duration is increased by 1 second.",
	},
	"ArtanisShieldOverloadShieldSurge": {
		Name: "Shield Surge",
		Text: "Increases Shield Overload's Shields by 75% while below 25% Health.",
	},
	"ArtanisSpearofAdunPurifierBeamTargetPurified": {
		Name: "Target Purified",
		Text: "If the target of Purifier Beam dies, it automatically recasts on the nearest visible enemy Hero.",
	},
	"ArtanisSpearofAdunSuppressionPulseOrbitalBombardment": {
		Name: "Orbital Bombardment",
		Text: "Suppression Pulse gains an additional charge.  There is a 10 second cooldown between uses.",
	},
	"ArtanisTwinBladesAmateurOpponent": {
		Name: "Amateur Opponent",
		Text: "Twin Blades attacks deal 150% bonus damage to non-Heroes.",
	},
	"ArtanisTwinBladesPsionicWound": {
		Name: "Psionic Wound",
		Text: "Twin Blades final strike lowers a Hero's Armor by 25 for 2 seconds, increasing all damage taken by 25%.",
	},
	"ArtanisTwinBladesReactiveParry": {
		Name: "Reactive Parry",
		Text: "Using Twin Blades grants 50 Physical Armor against the next enemy Hero Basic Attack, reducing its damage by 50%.\nStores up to 2 charges.",
	},
	"ArtanisTwinBladesTitanKiller": {
		Name: "Titan Killer",
		Text: "Twin Blades attacks against Heroes deal an additional 2% of the target's maximum Health in damage.",
	},
	"ArtanisTwinBladesTripleStrike": {
		Name: "Triple Strike",
		Text: "Increases Twin Blades's number of Basic Attacks to 3, but increase its cooldown by 1 second.",
	},
	"ArtanisTwinBladesZealotCharge": {
		Name: "Zealot Charge",
		Text: "Increase Twin Blades' charge distance by 100%.",
	},
	"ArthasAntiMagicShell": {
		Name: "Anti-Magic Shell",
		Text: "Activate to make Arthas immune to Spell Damage for 3 seconds and heal Arthas for 25% of the damage prevented.",
	},
	"ArthasDeathlord": {
		Name: "Deathlord",
		Text: "Reduces Death Coil's cooldown by 3 seconds and increase its range by 30%.",
	},
	"ArthasDeathsAdvance": {
		Name: "Death's Advance",
		Text: "Activate to increase Movement Speed by 30% for 3 seconds.\nPassive: Increases Movement Speed by 10%.",
	},
	"ArthasHeroicAbilityArmyoftheDead": {
		Name: "Army of the Dead",
		Text: "Summons Ghouls that last 15 seconds. Sacrifice Ghouls to heal for 267 Health.",
	},
	"ArthasHeroicAbilitySummonSindragosa": {
		Name: "Summon Sindragosa",
		Text: "Deals 230 damage and slows enemies by 60% for 4 seconds.  Also disables non-Heroic enemies and Structures for 20 seconds.",
	},
	"ArthasIceboundFortitude": {
		Name: "Icebound Fortitude",
		Text: "Activate to gain 25 Armor, reducing damage taken by 25%, and reduce the duration of Stuns, Slows, and Roots against Arthas by 75% for 3 seconds.",
	},
	"ArthasIcyTalons": {
		Name: "Icy Talons",
		Text: "Gain 3% Attack Speed for 1.5 seconds every time a Hero is damaged by Frozen Tempest, to a maximum of 60%. ",
	},
	"ArthasMasteryAbsoluteZeroSummonSindragosa": {
		Name: "Absolute Zero",
		Text: "Sindragosa flies twice as far.  Enemy Heroes are rooted for 2 seconds, and then slowed by 60% for 3.5 seconds.",
	},
	"ArthasMasteryBitingColdFrozenTempest": {
		Name: "Biting Cold",
		Text: "Each second an enemy is damaged by Frozen Tempest, it deals 12.5% bonus damage, up to a 50% bonus.",
	},
	"ArthasMasteryEmbraceDeathDeathCoil": {
		Name: "Embrace Death",
		Text: "Death Coil deals more damage and heals more the lower Arthas's current health is, to a maximum of 100% bonus damage and healing.",
	},
	"ArthasMasteryEternalHungerFrostmourneHungers": {
		Name: "Eternal Hunger",
		Text: "Quest: Use Frostmourne Hungers on an enemy Hero.\nReward: Increases the Mana it restores by 4, to a maximum of 40, and its damage by 4.",
	},
	"ArthasMasteryFrostPresenceHowlingBlast": {
		Name: "Frost Presence",
		Text: "Quest: Root enemy Heroes with Howling Blast.\nReward: After rooting 5 Heroes, Howling Blast's cooldown is reduced by 2 seconds.\nReward: After rooting 10 Heroes, Howling Blast's range is increased by 30%.\nReward: After rooting 20 Heroes, Howling Blast also roots enemies in its path.",
	},
	"ArthasMasteryFrostStrikeFrostmourneHungers": {
		Name: "Frost Strike",
		Text: "Reduces Frostmourne Hungers' cooldown by 2 seconds. Frostmourne Hungers also slows the enemy by 50% for 1.5 seconds.",
	},
	"ArthasMasteryFrostmourneFeedsFrostmourneHungers": {
		Name: "Frostmourne Feeds",
		Text: "Increases the amount of Basic Attacks empowered by Frostmourne Hungers to 2.",
	},
	"ArthasMasteryFrozenWastesFrozenTempest": {
		Name: "Frozen Wastes",
		Text: "Frozen Tempest Mana cost reduced by 4 per second.\nQuest: Damage enemy Heroes with Frozen Tempest.\nReward: After damaging enemy Heroes 150 times with Frozen Tempest, the Movement and Attack Speed slows of Frozen Tempest last an extra 1.5 seconds against enemy Heroes.",
	},
	"ArthasMasteryImmortalCoilDeathCoil": {
		Name: "Immortal Coil",
		Text: "Gain the healing effect of Death Coil even when used on enemies. If Death Coil is used on Arthas, it heals for an additional 50% bonus healing.",
	},
	"ArthasMasteryLegionOfNorthrendArmyoftheDead": {
		Name: "Legion of Northrend",
		Text: "3 additional Ghouls are created.  Ghouls heal for an additional 50% and last 5 seconds longer. ",
	},
	"ArthasMasteryRemorselessWinterFrozenTempest": {
		Name: "Remorseless Winter",
		Text: "Enemy Heroes that remain within Frozen Tempest for 3 seconds are rooted for 1.2 seconds. This effect can only happen once every 10 seconds.",
	},
	"ArthasRime": {
		Name: "Rime",
		Text: "Every 5 seconds, gain 60 Physical Armor against the next enemy Hero Basic Attack, reducing the damage taken by 60%. Stores up to 3 charges.",
	},
	"ArthasRuneTap": {
		Name: "Rune Tap",
		Text: "Every 3rd Basic Attack heals Arthas for 3.9% of his max Health.",
	},
	"ArthasShatteredArmor": {
		Name: "Shattered Armor",
		Text: "Enemy Heroes hit by Howling Blast have their Armor reduced by 15 for 4 seconds.",
	},
	"ArtifactAbilityPower": {
		Name: "Brilliant Topaz",
		Text: "Increases ability damage and healing",
	},
	"ArtifactAttackDamage": {
		Name: "Primal Ruby",
		Text: "Increases basic attack damage",
	},
	"ArtifactAttackSpeed": {
		Name: "Skyfire Emerald",
		Text: "Increases speed of basic attacks",
	},
	"ArtifactCooldownReduction": {
		Name: "Royal Diamond",
		Text: "Reduces cooldowns for abilities",
	},
	"ArtifactHealthRegen": {
		Name: "Warrior's Resolve",
		Text: "Increases Health regeneration",
	},
	"ArtifactLifesteal": {
		Name: "Blood Siphon",
		Text: "Basic attacks restore Health",
	},
	"ArtifactManaRegen": {
		Name: "Mana Infusion",
		Text: "Increases Mana regeneration",
	},
	"ArtifactMaxHealth": {
		Name: "Bold Amethyst",
		Text: "Increases maximum Health",
	},
	"ArtifactMaxMana": {
		Name: "Khaydarin Amulet",
		Text: "Increases maximum Mana",
	},
	"ArtifactMovementSpeed": {
		Name: "Wirt's Leg",
		Text: "Increases movement speed",
	},
	"ArtifactShields": {
		Name: "Nexus Sapphire",
		Text: "Grants shields",
	},
	"ArtifactSiegeDamage": {
		Name: "Destroyer",
		Text: "Increases damage vs. minions and towers",
	},
	"AurielAngelicFlightTalent": {
		Name: "Angelic Flight",
		Text: "After 2 seconds, fly to a target location.",
	},
	"AurielBlindingFlashSacredSweepTalent": {
		Name: "Blinding Flash",
		Text: "Enemies hit by the center area of Sacred Sweep are blinded for 2.2 seconds.",
	},
	"AurielBurstingLightRayOfHeavenTalent": {
		Name: "Bursting Light",
		Text: "Reduces the cooldown of Ray of Heaven by 2 seconds.",
	},
	"AurielConvergingForceSacredSweepTalent": {
		Name: "Converging Force",
		Text: "Enemies hit by the outer area are pushed slightly toward the center.",
	},
	"AurielDiamondResolveCrystalAegisTalent": {
		Name: "Diamond Resolve",
		Text: "When Crystal Aegis expires, it grants the target 50 Armor for 5 seconds, reducing damage taken by 50%.",
	},
	"AurielEmpathicLinkBestowHopeTalent": {
		Name: "Empathic Link",
		Text: "Auriel stores 25% of damage taken by allies with Bestow Hope.",
	},
	"AurielEnergizedCordRayOfHeavenTalent": {
		Name: "Energized Cord",
		Text: "Increases the energy stored from Auriel's Basic Attacks to 100% of the damage against Heroes and 50% of the damage against non-Heroes.\nDoes not affect Auriel's Bestow Hope ally.",
	},
	"AurielGlimmerofHopeRayOfHeavenTalent": {
		Name: "Glimmer of Hope",
		Text: "Collecting a Regeneration Globe reduces the cost of Auriel's next Ray of Heaven by 75%.",
	},
	"AurielHeavyBurdenDetainmentStrikeTalent": {
		Name: "Heavy Burden",
		Text: "When an enemy Hero is stunned by Detainment Strike, their Movement Speed is slowed by 40% for 3 seconds after the stun.",
	},
	"AurielHeroicCrystalAegis": {
		Name: "Crystal Aegis",
		Text: "Place an allied Hero into Stasis for 2 seconds. Upon expiration, Crystal Aegis deals 270 damage to all nearby enemies.",
	},
	"AurielHeroicResurrect": {
		Name: "Resurrect",
		Text: "Channel on the spirit of a dead ally for 0.5 seconds. After a 5 second delay, they are brought back to life with 50% of their maximum Health at the location where they died.",
	},
	"AurielIncreasingClaritySacredSweepTalent": {
		Name: "Increasing Clarity",
		Text: "Quest: Every time Sacred Sweep hits a Hero in the center, increase the center damage by 2, up to 50. \nReward: After hitting 25 Heroes, this center damage bonus is increased to 150.",
	},
	"AurielLightSpeedResurrectTalent": {
		Name: "Light Speed",
		Text: "Resurrected allies gain 199.7% increased Movement Speed, decaying over 4 seconds.  While a resurrected ally remains alive, Resurrect's next cooldown recharges 100% faster.",
	},
	"AurielMajesticSpanSacredSweepTalent": {
		Name: "Majestic Span",
		Text: "Increases the radius of Sacred Sweep by 15%.",
	},
	"AurielPiercingLashDetainmentStrikeTalent": {
		Name: "Piercing Lash",
		Text: "Detainment Strike now pierces and hits all enemy Heroes in a line, reducing the cooldown by 2 seconds for each Hero hit.",
	},
	"AurielRepeatedOffenseDetainmentStrikeTalent": {
		Name: "Repeated Offense",
		Text: "Quest: Every time Detainment Strike stuns a Hero, increase the stun damage by 10, up to 60.\nReward: After stunning 6 Heroes, increase this damage bonus to 250.",
	},
	"AurielRepellingStrikeDetainmentStrikeTalent": {
		Name: "Repelling Strike",
		Text: "Enemies hit by Detainment Strike are knocked back 35% farther.",
	},
	"AurielReservoirofHopeRayOfHeavenTalent": {
		Name: "Reservoir of Hope",
		Text: "Quest: Each maximum energy Ray of Heaven Auriel casts increases the maximum amount of energy that can be stored by 75.",
	},
	"AurielRighteousAssaultSacredSweepTalent": {
		Name: "Righteous Assault",
		Text: "Reduces the cooldown of Sacred Sweep by 3 seconds for each enemy Hero hit by its center.",
	},
	"AurielSearingLightRayOfHeavenTalent": {
		Name: "Searing Light",
		Text: "Ray of Heaven also deals damage to enemies in the area equal to 30% of the energy consumed.",
	},
	"AurielShieldOfHopeTalent": {
		Name: "Shield of Hope",
		Text: "Activate to grant all nearby allied Heroes a shield for 3 seconds equal to 50% of the amount of Health they are missing.",
	},
	"AurielSwiftSweepSacredSweepTalent": {
		Name: "Swift Sweep",
		Text: "Increases the cast speed of Sacred Sweep by 50%.",
	},
	"AurielWillofHeavenBestowHopeTalent": {
		Name: "Will of Heaven",
		Text: "Allies with Bestow Hope gain 20% Attack Speed.",
	},
	"AurielWrathofHeavenBestowHopeTalent": {
		Name: "Wrath of Heaven",
		Text: "Allies with Bestow Hope gain 10% Spell Power.",
	},
	"AzmodanBloodForBlood": {
		Name: "Blood for Blood",
		Text: "Activate to deal 10% of target enemy Hero's Max Health and heal for twice that amount.",
	},
	"AzmodanBoltoftheStorms": {
		Name: "Bolt of the Storm",
		Text: "Activate to teleport to a nearby location.",
	},
	"AzmodanBombardment": {
		Name: "Bombardment",
		Text: "[PH]Reward: Killing an enemy Hero within the last 50% of Globe of Annihilation's range permanently decreases Globe of Annihilation's cooldown by 5 seconds.",
	},
	"AzmodanBribe": {
		Name: "Bribe",
		Text: "Kill Minions to bribe a Mercenary",
	},
	"AzmodanGluttonousWard": {
		Name: "Gluttonous Ward",
		Text: "Activate to place a ward on the ground that restores 3.3% of Azmodan's maximum Health and Mana every second for 10 seconds. Can be cast while channeling All Shall Burn.",
	},
	"AzmodanHellishHirelings": {
		Name: "Hellish Hirelings",
		Text: "The damage bonus from General of Hell is doubled on Mercenaries.",
	},
	"AzmodanHeroicAbilityBlackPool": {
		Name: "Black Pool",
		Text: "Create a pool that empowers Azmodan, his Demons, and allied Minions, increasing their Basic Attack and Ability damage by 75%. Pools last 5 seconds.\nStores up to 2 charges.",
	},
	"AzmodanHeroicAbilityDemonicInvasion": {
		Name: "Demonic Invasion",
		Text: "Rain a small army of Demonic Grunts down on enemies, dealing 104 damage per impact. Demon Grunts deal 51 damage and have 365 health.",
	},
	"AzmodanMasteryArmyOfHell": {
		Name: "Army of Hell",
		Text: "Increases Demon Warrior's damage by 20% and reduces their Mana cost by 20.",
	},
	"AzmodanMasteryBattleborn": {
		Name: "Battleborn",
		Text: "If Globe of Annihilation hits an enemy, a Demon Warrior is summoned at the impact point.",
	},
	"AzmodanMasteryBlackPoolFifthCircle": {
		Name: "Fifth Circle",
		Text: "Black Pool makes Basic Attacks and Abilities slow enemies Attack and Movement Speeds by 40% for 3 seconds.",
	},
	"AzmodanMasteryBoundMinion": {
		Name: "Bound Minion",
		Text: "Using General of Hell on a Minion grants them 75 Armor from non-Heroic sources, decreasing the damage they take by 75%, while also  increasing the damage they deal to non-Heroic targets by 100%. Lasts 30 seconds.",
	},
	"AzmodanMasteryDemonicSmite": {
		Name: "Demonic Smite",
		Text: "Every 7.5 seconds, the Demon Lieutenant will blast a non-Heroic unit dealing 411 damage.",
	},
	"AzmodanMasteryEnduringWarriors": {
		Name: "Enduring Warriors",
		Text: "Demons last 5 seconds longer and do an additional 20% attack damage.",
	},
	"AzmodanMasteryForcedRecruitment": {
		Name: "Forced Recruitment",
		Text: "Reduces General of Hell's cooldown by 10 seconds and gains an additional charge.",
	},
	"AzmodanMasteryGluttony": {
		Name: "Gluttony",
		Text: "All Shall Burn heals Azmodan for 25% of the damage dealt.",
	},
	"AzmodanMasteryHedonism": {
		Name: "Hedonism",
		Text: "Reduces Globe of Annihilation's Mana cost by 30.",
	},
	"AzmodanMasteryHellforgedArmor": {
		Name: "Hellforged Armor",
		Text: "Demon Warriors deal 10 damage to nearby enemies every second and are granted 50 Armor against non-Heroic sources, decreasing the damage they take 50%.",
	},
	"AzmodanMasteryInfernalGlobe": {
		Name: "Infernal Globe",
		Text: "Globe of Annihilation's cast time is decreased by 40% and its missile speed is increased by 40%.",
	},
	"AzmodanMasteryInfusedPower": {
		Name: "Infused Power",
		Text: "All Shall Burn's damage growth maximum is increased by 40%.",
	},
	"AzmodanMasteryMarchOfSin": {
		Name: "March of Sin",
		Text: "Increases Azmodan's Movement Speed while channeling All Shall Burn to 80% of normal speed.",
	},
	"AzmodanMasteryMasterOfDestruction": {
		Name: "Master of Destruction",
		Text: "Reduces the Mana cost from 16 to 10 per second. Every 1.5 seconds spent channeling increases the range by 8.5%, to a maximum of 25.4%.",
	},
	"AzmodanMasteryPerishingFlame": {
		Name: "Perishing Flame",
		Text: "When the Grunts die they explode, dealing 68 damage to nearby enemies. Deals 100% increased damage to Minions, Mercenaries, Monsters, and Structures.",
	},
	"AzmodanMasterySiegingWrath": {
		Name: "Sieging Wrath",
		Text: "Quest: Every time Globe of Annihilation hits a Hero, its damage is increased by 3, up to 90.\nReward: After hitting 30 Heroes, the range of Globe of Annihilation is increased by 33% and its damage is increased by an additional 100.",
	},
	"AzmodanMasteryTasteForBlood": {
		Name: "Taste for Blood",
		Text: "Quest: Kill enemy Minions and Heroes within 1.5 seconds of being hit by Globe of Annihilation.\nReward: Enemy Minions that die increase its damage by 2.\nReward: Enemy Heroes that die increase its damage by 10.\nMaximum bonus damage of 500.",
	},
	"AzmodanSinforSin": {
		Name: "Sin for Sin",
		Text: "Activate to deal 10% of target enemy Hero's Max Health and heal for twice that amount. Can be cast while channeling All Shall Burn.",
	},
	"AzmodanSinsGrasp": {
		Name: "Sin's Grasp",
		Text: "Activate to curse an enemy Hero, dealing 247.5 damage over 8 seconds. Minion kills reduce this cooldown by 10 seconds. Can be cast while channeling All Shall Burn.",
	},
	"BarbarianAmplifiedHealing": {
		Name: "Amplified Healing",
		Text: "Increases regeneration effects and all healing received by 30%.",
	},
	"BarbarianCombatStyleEndlessFury": {
		Name: "Endless Fury",
		Text: "Increases maximum Fury to 200.",
	},
	"BarbarianCombatStyleFerociousHealing": {
		Name: "Ferocious Healing",
		Text: "Consume 20 Fury to heal 10.2% of Sonya's maximum Health.\nUsable while Whirlwinding.",
	},
	"BarbarianCombatStyleNoEscape": {
		Name: "No Escape",
		Text: "Increases the Movement Speed bonus from using Basic and Heroic Abilities to 25%.",
	},
	"BarbarianCombatStyleShotofFury": {
		Name: "Shot of Fury",
		Text: "Increases maximum Fury by 50. Sonya's Trait can be activated to gain 50 Fury. \nUsable while Whirlwinding.",
	},
	"BarbarianHeroicAbilityLeap": {
		Name: "Leap",
		Text: "Leap into the air, dealing 135 damage to nearby enemies, and stunning them for 1.2 seconds.",
	},
	"BarbarianHeroicAbilityWrathoftheBerserker": {
		Name: "Wrath of the Berserker",
		Text: "Increase damage dealt by 40%. Reduce the duration of stuns, slows, and roots against Sonya by 50%. Lasts 15 seconds, and extends by 1 second for every 10 Fury gained.",
	},
	"BarbarianMasteryAftershock": {
		Name: "Aftershock",
		Text: "Using Seismic Slam reduces its Fury cost by 50% for 2 seconds.",
	},
	"BarbarianMasteryAngerManagementWrathoftheBerserker": {
		Name: "Anger Management",
		Text: "Increases all Fury generated during Wrath of the Berserker by 100%.",
	},
	"BarbarianMasteryArreatCraterLeap": {
		Name: "Arreat Crater",
		Text: "Leap leaves behind an impassable crater for 5 seconds.  Reduces Leap's cooldown by 20 seconds.",
	},
	"BarbarianMasteryBoonOfTheAncients": {
		Name: "Boon Of The Ancients",
		Text: "Boon Of The Ancients \nHitting an enemy reduces Ancient Spear's cooldown by 5 seconds.",
	},
	"BarbarianMasteryCompositeSpearAncientSpear": {
		Name: "Composite Spear",
		Text: "Increases the range of Ancient Spear by 30% and increases Fury generated by 20.",
	},
	"BarbarianMasteryCycloneWhirlwind": {
		Name: "Cyclone",
		Text: "Increases the radius of Whirlwind by 80%.",
	},
	"BarbarianMasteryDustDevils": {
		Name: "Dust Devils",
		Text: "Whirlwind creates Tornadoes that do an additional 25% damage.",
	},
	"BarbarianMasteryEnduringWhirlwind": {
		Name: "Enduring Whirlwind",
		Text: "Whirlwind duration increased by 50%.",
	},
	"BarbarianMasteryFuriousBlowSeismicSlam": {
		Name: "Furious Blow",
		Text: "Increases Seismic Slam damage by 50%, but costs 40 Fury.",
	},
	"BarbarianMasteryHamstringSpiralWhirlwind": {
		Name: "Hamstring Spiral",
		Text: "Whirlwind slows enemy Movement Speed by 25% for 1 second.",
	},
	"BarbarianMasteryHurricaneWhirlwind": {
		Name: "Hurricane",
		Text: "Reduces the cooldown of Whirlwind by 1 second and Whirlwind removes all slows and roots on Sonya.",
	},
	"BarbarianMasteryLifeFunnelWhirlwind": {
		Name: "Life Funnel",
		Text: "Increases the healing of Whirlwind to 35%, tripled versus Heroes.",
	},
	"BarbarianMasteryMysticalSpearAncientSpear": {
		Name: "Mystical Spear",
		Text: "Reduces the cooldown of Ancient Spear by 3 seconds. Sonya is pulled to the target location even if an enemy isn't hit.",
	},
	"BarbarianMasteryPoisonedSpearAncientSpear": {
		Name: "Poisoned Spear",
		Text: "Ancient Spear deals an additional 150% damage over 6 seconds.",
	},
	"BarbarianMasteryShatteredGroundSeismicSlam": {
		Name: "Shattered Ground",
		Text: "Increases Seismic Slam splash damage to 75% of primary target damage.",
	},
	"BarbarianMasteryStrengthFromEarthSeismicSlam": {
		Name: "Strength from Earth",
		Text: "Strength From Earth \nSeismic Slam heals for 100% of damage dealt.",
	},
	"BarbarianMasteryWarPaint": {
		Name: "War Paint",
		Text: "Basic Attacks heal Sonya for 30% of the damage dealt.",
	},
	"BattleMomentumCrusader": {
		Name: "Blessed Momentum",
		Text: "Basic Attacks reduce Johanna's Basic Ability cooldowns by 0.5 seconds.",
	},
	"BattleMomentumETC": {
		Name: "Battle Momentum",
		Text: "Basic Attacks reduce Ability cooldowns by 0.5 seconds.",
	},
	"BattleMomentumKerrigan": {
		Name: "Bladed Momentum",
		Text: "Basic Attacks reduce Kerrigan's Basic Ability cooldowns by 0.6 seconds.",
	},
	"BattleMomentumMuradin": {
		Name: "Iron-forged Momentum",
		Text: "Basic Attacks reduce the cooldown of Storm Bolt and Thunderclap by -0 seconds.",
	},
	"BattleMomentumNova": {
		Name: "Battle Momentum",
		Text: "Basic Attacks reduce Ability cooldowns by 0.5 seconds.",
	},
	"BattleMomentumRehgar": {
		Name: "Battle Momentum",
		Text: "Basic Attacks reduce Ability cooldowns by 0.5 seconds.",
	},
	"BattleMomentumSylvanas": {
		Name: "Battle Momentum",
		Text: "Basic Attacks reduce Ability cooldowns by 0.5 seconds.",
	},
	"BattleMomentumTyrael": {
		Name: "Angelic Momentum",
		Text: "Basic Attacks reduce Tyrael's Basic Ability cooldowns by 0.8 seconds.",
	},
	"BrightwingArcaneBarrageArcaneFlare": {
		Name: "Arcane Barrage",
		Text: "Increases the range of Arcane Flare by 50%.",
	},
	"BrightwingBouncyDustPixieDust": {
		Name: "Bouncy Dust",
		Text: "Pixie Dust bounces to another nearby ally upon impact.",
	},
	"BrightwingDoubleWyrmholeBlinkHeal": {
		Name: "Double Wyrmhole",
		Text: "Blink Heal can be cast a second time on a different target within 2 seconds without consuming a charge.",
	},
	"BrightwingDreamShotArcaneFlare": {
		Name: "Dream Shot",
		Text: "Increase the range of Arcane Flare by 50%. If an enemy Hero is hit by the inner radius, reduce its cooldown to 1 second.",
	},
	"BrightwingGreaterPolymorphPolymorph": {
		Name: "Greater Polymorph",
		Text: "Increases the duration of Polymorph by 0.8 seconds.",
	},
	"BrightwingHyperShiftPhaseShift": {
		Name: "Hyper Shift",
		Text: "Each allied Hero healed by Soothing Mist reduces the cooldown of Phase Shift by 3 seconds.",
	},
	"BrightwingManicPixiePixieDust": {
		Name: "Manic Pixie",
		Text: "When Soothing Mist heals an ally with Pixie Dust on them, they heal for an additional 104 over 4 seconds.",
	},
	"BrightwingMistifiedSoothingMist": {
		Name: "Mistified",
		Text: "Reduces the cooldown of Soothing Mist by 0.5 seconds every time Brightwing casts a Basic Ability.",
	},
	"BrightwingPeekabooPhaseShift": {
		Name: "Peekaboo!",
		Text: "Phase Shifting to an ally casts a free Pixie Dust on them and reveals a large area around them and all enemies in it for 6 seconds. ",
	},
	"BrightwingPixieBoostPixieDust": {
		Name: "Pixie Boost",
		Text: "Pixie Dust gives 50% bonus Move Speed, decaying to 20% over 1.5 seconds.",
	},
	"BrightwingPixieCharm": {
		Name: "Pixie Charm",
		Text: "Each time Brightwing heals an Allied Hero with Soothing Mist, she gains a stack of Bribe. Use 20 stacks to bribe a Mercenary, instantly defeating them. Does not work on Bosses. Maximum of 80 stacks.",
	},
	"BrightwingRevitalizingMistSoothingMist": {
		Name: "Revitalizing Mist",
		Text: "Healing another Hero with Soothing Mist increases their healing received from Soothing Mist by 25% for 6 seconds. Stacks 3 times.",
	},
	"BrightwingUnfurlingSpraySoothingMist": {
		Name: "Unfurling Spray",
		Text: "Increases Soothing Mist's range by 20%.",
	},
	"BrightwingUnstableAnomalyPolymorph": {
		Name: "Unstable Anomaly",
		Text: "When Polymorph ends, deal 114 damage to the target and all nearby enemies.",
	},
	"BrightwingWoundedAnimalPolymorph": {
		Name: "Wounded Animal",
		Text: "Increase the Movement Speed slow of Polymorph to 30%.",
	},
	"BucketCallOfTheStorm": {
		Name: "Onslaught of the Storm",
		Text: "Onslaught of the Storm\nEnemies killed by Abilities explode for 500 damage in a small area.",
	},
	"BucketConstantHeroism": {
		Name: "Constant Heroism",
		Text: "Constant Heroism \nYour Heroic Ability's cooldown is reduced by 30%.",
	},
	"BucketFuryOfTheStorm": {
		Name: "Fury of the Storm",
		Text: "Fury of the Storm\nBasic Attacks bounce twice to nearby enemies for 50% damage.",
	},
	"ButcherCleaver": {
		Name: "Cleaver",
		Text: "Basic Attacks deal 35% of The Butcher's Basic Attack Damage in an area around the target.",
	},
	"ButcherHeroicAbilityButcherFurnaceBlast": {
		Name: "Furnace Blast",
		Text: "After a 3 second delay, fire explodes around The Butcher dealing 500 damage to enemies.\nCan be cast while using Ruthless Onslaught.",
	},
	"ButcherHeroicAbilityLambToTheSlaughter": {
		Name: "Lamb to the Slaughter",
		Text: "Throw a hitching post that attaches to the nearest enemy Hero after a 1 second delay. This deals 171 damage and causes the enemy to be chained to the post and Silenced for 3 seconds.",
	},
	"ButcherMasteryButchersBrandCraveFlesh": {
		Name: "Crave Flesh",
		Text: "While an enemy is affected by Butcher's Brand, gain 30% Movement Speed.",
	},
	"ButcherMasteryButchersBrandExposeFlesh": {
		Name: "Expose Flesh",
		Text: "Expose Flesh\nWhile a target is under the effects of Butcher's Brand, they are also Vulnerable, increasing all damage taken by 25%.",
	},
	"ButcherMasteryButchersBrandInsatiableBlade": {
		Name: "Insatiable Blade",
		Text: "While facing a branded enemy, The Butcher's Movement Speed is increased by 25%.",
	},
	"ButcherMasteryFiresofHell": {
		Name: "Fires of Hell",
		Text: "Furnace Blast explodes a second time 3 seconds after the initial explosion.",
	},
	"ButcherMasteryFreshMeatBloodFrenzy": {
		Name: "Blood Frenzy",
		Text: "Basic Attacks against enemy Heroes grant 5% increased Attack and Movement Speed for 3 seconds, up to 25%.",
	},
	"ButcherMasteryFreshMeatVictuals": {
		Name: "Victuals",
		Text: "Every time a nearby enemy minion dies, The Butcher heals 5% of his maximum health.",
	},
	"ButcherMasteryHamstringBrutalStrike": {
		Name: "Brutal Strike",
		Text: "After using Hamstring, The Butcher's next 3 Basic Attacks within 5 seconds deal an additional 15% damage.",
	},
	"ButcherMasteryHamstringCheapShot": {
		Name: "Cheap Shot",
		Text: "Hamstring does 100% more damage to targets affected by a Slow, Root, or Stun. ",
	},
	"ButcherMasteryHamstringChopMeat": {
		Name: "Chop Meat",
		Text: "Increases Hamstring damage by 100% to Minions and Mercenaries.",
	},
	"ButcherMasteryHamstringCripplingSlam": {
		Name: "Crippling Slam",
		Text: "Hamstring's slow no longer fades out, and the duration is increased by 30%.",
	},
	"ButcherMasteryHamstringFlailAxe": {
		Name: "Flail Axe",
		Text: "Increases the length of Hamstring by 40%.",
	},
	"ButcherMasteryHamstringInvigoration": {
		Name: "Invigoration",
		Text: "If Hamstring hits a Hero, half of the Mana cost is refunded and the cooldown is reduced by 1 second.",
	},
	"ButcherMasteryRuthlessOnslaughtSavageCharge": {
		Name: "Savage Charge",
		Text: "Ruthless Onslaught deals bonus damage to Heroes equal to 10% of their maximum Health.",
	},
	"ButcherMasteryRuthlessOnslaughtUnrelentingPursuit": {
		Name: "Unrelenting Pursuit",
		Text: "Reduces the cooldown of Ruthless Onslaught by 33.3% upon impact.",
	},
	"ButcherMasterySlaughterhouse": {
		Name: "Slaughterhouse",
		Text: "Lamb to the Slaughter now chains all enemy Heroes in range.",
	},
	"ButcherMeatShield": {
		Name: "Meat Shield",
		Text: "When Ruthless Onslaught impacts an enemy Hero, The Butcher gains 50 Spell Armor for 2.5 seconds, reducing Ability Damage taken by 50%.",
	},
	"ButcherTalentEnraged": {
		Name: "Enraged",
		Text: "While below 50% of his maximum Health, taking damage causes The Butcher to become Enraged for 10 seconds, gaining 40% Attack Speed and 15 Armor, reducing damage taken by 15%.\nThis effect has a 25 second cooldown.",
	},
	"ChenAccumulatingFlame": {
		Name: "Accumulating Flame",
		Text: "Quest: Every time a Hero is Ignited with Breath of Fire, increase its damage over time by 2%, up to 60%.\nReward: After Igniting 30 Heroes, gain 25 Armor for 1 second per Hero Ignited, reducing damage taken by 25%.",
	},
	"ChenAnotherRound": {
		Name: "Another Round",
		Text: "Increases the radius of Keg Smash by 50%. After hitting an enemy Hero with Keg Smash, the cooldown of the next Basic Ability used will be reduced by 3 seconds.",
	},
	"ChenBolderFlavor": {
		Name: "Bolder Flavor",
		Text: "Fortifying Brew instantly Shields Chen for 318. Additionally, Shields persist for 1 extra second after he stops drinking.",
	},
	"ChenBrewmastersBalance": {
		Name: "Brewmaster's Balance",
		Text: "While at or below 50 Brew, gain 20% Movement Speed. While at or above 50 Brew, regenerate an additional 18 Health per second.",
	},
	"ChenDeadlyStrike": {
		Name: "Deadly Strike",
		Text: "Flying Kick no longer costs Brew. Additionally, its damage is increased by 100% while Chen has Shields from Fortifying Brew.",
	},
	"ChenElementalConduit": {
		Name: "Elemental Conduit",
		Text: "Increases the spirits' Health to 100% of Chen's maximum Health and empowers their abilities.\nStorm can grant the spirits Unstoppable. \nEarth's leap cooldown is reduced by 3 seconds. \nFire grants Attack Speed as long as it is alive.",
	},
	"ChenElusiveBrawler": {
		Name: "Elusive Brawler",
		Text: "Activate to evade enemy Basic Attacks for 2 seconds. Chen's Basic Attacks reduce this cooldown by 2 seconds.",
	},
	"ChenFlyingLeap": {
		Name: "Flying Leap",
		Text: "Increases Flying Kick's range by 15%.",
	},
	"ChenFreshestIngredients": {
		Name: "Freshest Ingredients",
		Text: "Quest: Gathering a Regeneration Globe increases Chen's Health Regeneration by 1 per second, up to 30.\nReward: After gathering 30 Regeneration Globes, 50% of the Shield from Fortifying Brew persists indefinitely after drinking.",
	},
	"ChenGroundingBrew": {
		Name: "Grounding Brew",
		Text: "Fortifying Brew grants 30 Spell Armor while drinking, reducing Ability Damage taken by 30%.",
	},
	"ChenHeroicAbilityStormEarthFire": {
		Name: "Storm, Earth, Fire",
		Text: "After 1 second, Chen splits into three elemental spirits for 12.2 seconds, each with 75% of Chen's maximum Health and a unique ability. If all three spirits are killed, Chen will die as well.\nStorm can grant the spirits Movement Speed. \nEarth can leap to an area and slow enemies. \nFire can grant the spirits Attack Speed.",
	},
	"ChenHeroicAbilityWanderingKeg": {
		Name: "Wandering Keg",
		Text: "Roll around inside an Unstoppable barrel with 60% increased Movement Speed, dealing 59 damage to enemies in the way and knocking them back. Lasts for 5 seconds.",
	},
	"ChenKegToss": {
		Name: "Keg Toss",
		Text: "Quest: Every time Keg Smash hits a Hero, increase its damage by 3, up to 60. \nReward: After hitting 20 Heroes, increase its range by 125% and gain an additional charge.",
	},
	"ChenMasteryFortifyingBrewEnoughToShare": {
		Name: "Enough to Share",
		Text: "Fortifying Brew also Shields nearby allied Heroes for 54 per second.",
	},
	"ChenMasteryKegSmashATouchOfHoney": {
		Name: "A Touch of Honey",
		Text: "Increase the slow from Keg Smash to 40%.",
	},
	"ChenMasteryWanderingKegUntappedPotential": {
		Name: "Untapped Potential",
		Text: "Wandering Keg increases Movement Speed by 60% and grants 50 Armor, reducing damage taken by 50%.",
	},
	"ChenPressurePoint": {
		Name: "Pressure Point",
		Text: "Flying Kick slows the target enemy by 35% for 1 second, or by 70% if they're soaked in Brew.",
	},
	"ChenPurifyingBrew": {
		Name: "Purifying Brew",
		Text: "The next Stun or Silence used against Chen has its duration reduced by 75% and resets the cooldown of Fortifying Brew. Can only trigger once every 15 seconds. While channeling Fortifying Brew gain 30 Spell Armor, reducing Ability Damage taken by 30%.",
	},
	"ChenRefreshingElixir": {
		Name: "Refreshing Elixir",
		Text: "Increases regeneration effects and all healing received by 30%, or by 60% while drinking Fortifying Brew.",
	},
	"ChenRingofFire": {
		Name: "Ring of Fire",
		Text: "After using Breath of Fire, ignite in a fiery aura, dealing 52 damage every second to nearby enemies for 3 seconds.",
	},
	"ChenStormstoutSecretRecipe": {
		Name: "Stormstout Secret Recipe",
		Text: "Heal for 1.5% of Chen's maximum Health every time his Basic Abilities hit an enemy Hero.",
	},
	"ChenWitheringFlames": {
		Name: "Withering Flames",
		Text: "Setting an enemy Hero on fire reduces their Spell Power by 25% for 3 seconds.",
	},
	"ChoCallousedHide": {
		Name: "Calloused Hide",
		Text: "Ogre Hide also increases healing received by 15%.",
	},
	"ChoConsumingFire": {
		Name: "Consuming Fire",
		Text: "Consuming Blaze heals for 150% more when a Hero is Ignited.",
	},
	"ChoCthunsGift": {
		Name: "C'Thun's Gift",
		Text: "Cho's Basic Attack becomes ranged and slows targets by 20% for 2 seconds.",
	},
	"ChoEnragedRegeneration": {
		Name: "Enraged Regeneration",
		Text: "While Gall's Ogre Rage is active, Cho's Health Regeneration is increased by 200%.",
	},
	"ChoFavorOfTheOldGods": {
		Name: "Favor of the Old Gods",
		Text: "Upheaval roots enemy Heroes for 1.5 seconds.",
	},
	"ChoFirestarter": {
		Name: "Firestarter",
		Text: "Basic Attacks against ignited Heroes decreases the cooldown of Consuming Blaze by 1 second.\nPassive: Reduce the cooldown of Consuming Blaze by 3 seconds. ",
	},
	"ChoFrenziedFists": {
		Name: "Frenzied Fists",
		Text: "Gain 75% Attack Speed for 5 seconds after using Surging Fist.",
	},
	"ChoFuelForTheFlame": {
		Name: "Fuel for the Flame",
		Text: "Quest: Every Minion killed near Cho increases the amount healed by Consuming Blaze by 0.3%.\nQuest: Every Hero Takedown increases the amount healed by Consuming Blaze by 3%.",
	},
	"ChoHeroicAbilityHammerOfTwilight": {
		Name: "Hammer of Twilight",
		Text: "Activate to swing the Hammer of Twilight, dealing 150 damage, pushing enemies away, and Stunning them for 0.8 seconds.\nPassive: Cho's Basic Attacks deal 100% increased damage.",
	},
	"ChoHeroicAbilityUpheaval": {
		Name: "Upheaval",
		Text: "After 1 second, pull enemies towards Cho'gall, slowing them by 25% for 3 seconds and dealing 175 damage.",
	},
	"ChoHourofTwilight": {
		Name: "Hour of Twilight",
		Text: "Decreases Cho'gall's death timer by 50%.",
	},
	"ChoPowerSurge": {
		Name: "Power Surge",
		Text: "Each Hero hit by Surging Fist reduces its cooldown by 4 seconds.",
	},
	"ChoRollback": {
		Name: "Rollback",
		Text: "Rune Bomb rolls back to Cho, damaging enemies in its path.",
	},
	"ChoRunedGauntlet": {
		Name: "Runed Gauntlet",
		Text: "Basic Attacks reduce the cooldown of Cho and Gall's Heroic Abilities by 0.8 seconds.",
	},
	"ChoRunicFeedback": {
		Name: "Runic Feedback",
		Text: "Gall's Runic Blast reduces Cho's Rune Bomb cooldown by 1 seconds per enemy hit, and 2 seconds for each Hero hit.",
	},
	"ChoSearedFlesh": {
		Name: "Seared Flesh",
		Text: "Each consecutive Basic Attack to an enemy Hero deals 20% more damage, to a maximum of 60% damage. This bonus lasts 5 seconds or until a different enemy is attacked.",
	},
	"ChoSurgingDash": {
		Name: "Surging Dash",
		Text: "While channeling Surging Fist, Cho is Unstoppable and heals for 150 Health per second.",
	},
	"ChoTalentMoltenBlock": {
		Name: "Molten Block",
		Text: "Activate to enter Stasis and gain Invulnerability for 3 seconds, damaging nearby enemies for 92 damage per second.",
	},
	"ChoTheWillofCho": {
		Name: "The Will of Cho",
		Text: "Takedowns permanently increase the Armor granted by Ogre Hide by 2, to a max of 50 extra Armor.",
	},
	"ChoTwilightVeil": {
		Name: "Twilight Veil",
		Text: "Activate to increase the Armor bonus of Ogre Hide by 300% for 2 seconds. If cast while Ogre Rage is active, it instantly swaps to Ogre Hide.",
	},
	"ChoUppercut": {
		Name: "Uppercut",
		Text: "Surging Fist deals additional damage to Heroes equal to 7% of their max Health.",
	},
	"ChromieBronzeTalons": {
		Name: "Bronze Talons",
		Text: "Increase Chromie's Basic Attack range by 35%. Using Sand Blast increases the damage of her next Basic Attack by 250%.",
	},
	"ChromieDragonsBreathDeepBreathing": {
		Name: "Deep Breathing",
		Text: "Quest: Every time Dragon's Breath hits a Hero increase its damage by 3, up to 60.\nReward: After hitting 20 Heroes, increase its damage by an additional 140 and also increase Chromie's sight radius by 100%.",
	},
	"ChromieDragonsBreathDragonsEye": {
		Name: "Dragon’s Eye",
		Text: "Dragon’s Breath deals 25% more damage to enemies in the center.",
	},
	"ChromieDragonsBreathEnvelopingAssault": {
		Name: "Enveloping Assault",
		Text: "Increase Dragon's Breath's radius by 25%.",
	},
	"ChromieDragonsBreathMobiusLoop": {
		Name: "Mobius Loop",
		Text: "Reduce Dragon’s Breath's Mana cost and cooldown by 50%, but also reduce its damage by 25%.",
	},
	"ChromieHereAndThere": {
		Name: "Here and There",
		Text: "Activate to swap Chromie's position with her Sand Blast Echo.",
	},
	"ChromieHeroicAbilitySlowingSands": {
		Name: "Slowing Sands",
		Text: "Place a sand vortex that greatly Slows enemies inside it. The longer it is active the more it Slows, up to 60% after 3 seconds.",
	},
	"ChromieHeroicAbilityTemporalLoop": {
		Name: "Temporal Loop",
		Text: "Choose an enemy Hero. After 3 seconds, they will teleport back to the location where Temporal Loop was cast on them.",
	},
	"ChromiePortBackToBaseByeBye": {
		Name: "Bye Bye!",
		Text: "Reduce the cast time of Hearthstone by 75%, and it is not interrupted by taking damage.",
	},
	"ChromieQuantumOverdrive": {
		Name: "Quantum Overdrive",
		Text: "Activate to increase Spell Power by 20% for 10 seconds.",
	},
	"ChromieReachingThroughTime": {
		Name: "Reaching through Time",
		Text: "Increases Sand Blast and Dragon’s Breath's range by 15%.",
	},
	"ChromieSandBlastFastForward": {
		Name: "Fast Forward",
		Text: "If Sand Blast travels at least 75% of its base distance and hits a Hero, its cooldown is reduced to 0.8 seconds.",
	},
	"ChromieSandBlastMountingSand": {
		Name: "Mounting Sand",
		Text: "Hitting 3 or more consecutive Sand Blasts without missing counts as hitting 3 Heroes.",
	},
	"ChromieSandBlastPiercingSands": {
		Name: "Piercing Sands",
		Text: "Sand Blast will now hit all Heroes in its path. Additional hits do not grant bonus damage or Quest progress.",
	},
	"ChromieSandBlastShiftingSands": {
		Name: "Shifting Sands",
		Text: "Hitting an enemy Hero with Sand Blast grants 8% Spell Power for 10 seconds, up to a maximum of 40%. The duration is refreshed whenever an enemy is damaged by Sand Blast.",
	},
	"ChromieSlowingSandsPocketOfTime": {
		Name: "Pocket of Time",
		Text: "Remove Slowing Sands Mana cost and increase the Slow from 60% to 80%.",
	},
	"ChromieTemporalLoopStuckInALoop": {
		Name: "Stuck in a Loop",
		Text: "Temporal Loop teleports the target a second time after the first.",
	},
	"ChromieTimeOut": {
		Name: "Time Out",
		Text: "Activate to place Chromie in Stasis and gain Invulnerability for up to 7 seconds. Can be reactivated to end the effect early.",
	},
	"ChromieTimeTrapAndorhalAnomaly": {
		Name: "Andorhal Anomaly",
		Text: "Time Trap gains 2 charges, and when the cooldown finishes, Chromie gains all charges at once. Maximum 3 active Time Traps.",
	},
	"ChromieTimeTrapChronoSickness": {
		Name: "Chrono Sickness",
		Text: "Reduce Time Trap's cooldown and Mana cost by 50%. After the Time Stop ends, the enemy is also Slowed by 50% for 4 seconds.",
	},
	"ChromieTimeTrapTimelySurprise": {
		Name: "Timely Surprise",
		Text: "Increase Time Trap's cast range by 33.3%. Additionally, the cooldowns of Sand Blast and Dragon's Breath are reset when Time Trap is triggered.",
	},
	"ChromieTimewalkersPursuit": {
		Name: "Timewalker's Pursuit",
		Text: "Activate to reveal the targeted area for 2 seconds. Enemies seen are revealed for 4 seconds.\nQuest: Gathering a Regeneration Globe increases Chromie's Mana Regeneration by 0.1, up to 1.5.\nReward: After gathering 15 Globes Chromie also gains 5% Spell Power.",
	},
	"CrusaderBlessedHammer": {
		Name: "Blessed Hammer",
		Text: "Activate to create 2 hammers that spiral outward from Johanna, dealing 84 damage to enemies hit.\nHitting Heroes with Shield Glare reduces the cooldown of this Ability by 8 seconds.",
	},
	"CrusaderBlindedByTheLight": {
		Name: "Blinded By The Light",
		Text: "Activate to grant nearby allied Heroes a Shield equal to 25% of their maximum Health for 3 seconds.\nHitting Heroes with Shield Glare reduces the cooldown of this Ability by 8 seconds.",
	},
	"CrusaderHeroicAbilityBlessedShield": {
		Name: "Blessed Shield",
		Text: "Deal 114 damage and Stun the first enemy hit for 1.5 seconds. Blessed Shield then bounces to 2 nearby enemies, dealing 57 damage and Stunning them for 0.8 seconds.",
	},
	"CrusaderHeroicAbilityFallingSword": {
		Name: "Falling Sword",
		Text: "Johanna leaps towards an area.  While in the air, she can steer the landing location by moving. \nAfter 2 seconds Johanna lands, dealing 183 damage to nearby enemies, Stunning them for 0.2 seconds, and Slowing them by 50% for 3 seconds.",
	},
	"CrusaderHeroicMasteryIndestructable": {
		Name: "Indestructible",
		Text: "Upon taking fatal damage, Johanna gains a Shield equal to her maximum Health for 4 seconds. This effect has a 120 second cooldown.",
	},
	"CrusaderHolyFury": {
		Name: "Holy Fury",
		Text: "Deal 12 damage per second to nearby enemies. Each Hero hit by Condemn increases this damage by 40% for 5 seconds.",
	},
	"CrusaderMasteryBlessedShieldRadiatingFaith": {
		Name: "Radiating Faith",
		Text: "Increases Blessed Shield's Stun duration by 33% and maximum enemies hit by 2.",
	},
	"CrusaderMasteryCondemnConviction": {
		Name: "Conviction",
		Text: "Movement Speed is increased by 25% while Condemn is charging up.",
	},
	"CrusaderMasteryCondemnEternalRetaliation": {
		Name: "Eternal Retaliation",
		Text: "Condemn's cooldown is lowered by 0.5 seconds for each enemy affected. Maximum of 10 targets.",
	},
	"CrusaderMasteryCondemnGravitationalPull": {
		Name: "Gravitational Pull",
		Text: "Gravitational Pull \nRange increased by 25%.",
	},
	"CrusaderMasteryCondemnKnightTakesPawn": {
		Name: "Knight Takes Pawn",
		Text: "Condemn deals 165 additional damage to Minions and Mercenaries and stuns them for 3 seconds.",
	},
	"CrusaderMasteryFallingSwordHeavensFury": {
		Name: "Heaven's Fury",
		Text: "While in the air, holy bolts rain down on enemies dealing 75 damage and reducing the cooldown of Falling Sword by 3 seconds for each enemy hit.",
	},
	"CrusaderMasteryIronSkinFanaticism": {
		Name: "Fanaticism",
		Text: "Increase the duration of Iron Skin by 2 seconds. While Iron Skin is active, Johanna gains 8% Movement Speed each time she takes damage, up to 40%.",
	},
	"CrusaderMasteryIronSkinHoldYourGround": {
		Name: "Hold Your Ground",
		Text: "Increases Iron Skin's Shield by 30% and its cooldown is reduced by 4 seconds.",
	},
	"CrusaderMasteryIronSkinReinforce": {
		Name: "Reinforce",
		Text: "Using a Basic Ability grants 50 Physical Armor against the next enemy Hero Basic Attack, reducing damage taken by 50%.\nStores up to 2 charges.",
	},
	"CrusaderMasteryIronSkinTheCrusadeMarchesOn": {
		Name: "The Crusade Marches On",
		Text: "Basic and Heroic Abilities lower the cooldown of Iron Skin by 1.5 seconds.",
	},
	"CrusaderMasteryLawsOfHope": {
		Name: "Laws of Hope",
		Text: "Activate to heal 30% of Johanna's max Health over 4 seconds.\nPassive: Regenerate 1.5 Health per second.",
	},
	"CrusaderMasteryPunishRighteousSmash": {
		Name: "Righteous Smash",
		Text: "Punish restores 10 Mana per enemy hit.",
	},
	"CrusaderMasteryPunishRoar": {
		Name: "Roar",
		Text: "Increase Punish's damage by 25%. This bonus is increased to 150% whenever Punish hits 2 or more Heroes.",
	},
	"CrusaderMasteryPunishSubdue": {
		Name: "Subdue",
		Text: "Hitting 2 or more Heroes at the same time with Punish increases the Slow to 80% and the Slow amount does not decay.\nQuest: Hit 4 or more Heroes at the same time with Punish.\nReward: Punish always Slows by 80% and the Slow amount no longer decays.",
	},
	"CrusaderMasteryShieldGlareHolyRenewal": {
		Name: "Holy Renewal",
		Text: "Every enemy Hero affected by Shield Glare restores 114 health.",
	},
	"CrusaderMasteryShieldGlarePathOfTheCrusade": {
		Name: "Path Of The Crusade",
		Text: "Path of the Crusade\nUsing Punish while Iron Skin is active refreshes the duration of Iron Skin.",
	},
	"CrusaderMasteryShieldGlareSinsExposed": {
		Name: "Sins Exposed",
		Text: "Shield Glare marks enemies for 4 seconds. The next time any ally damages them, they take 90 extra damage and the mark is removed.",
	},
	"CrusaderZealousGlare": {
		Name: "Zealous Glare",
		Text: "Increase the Blind duration of Shield Glare by 0.5 seconds. Johanna's Basic Attacks against Heroes with Shield Glare extend its duration by 1 second, up to a maximum of 3 seconds.",
	},
	"DVaBigShot": {
		Name: "Big Shot",
		Text: "Deal 250 damage to all enemies in a line. The cooldown of Call Mech is reduced by 8 seconds for each enemy Hero hit.\nRequires Pilot Mode.",
	},
	"DVaBigShotPewPewPew": {
		Name: "Pew! Pew! Pew!",
		Text: "Instead of a single shot, Big Shot fires 3 shots over 0.5 seconds. Each shot deals 50% damage.",
	},
	"DVaBoostersComingThrough": {
		Name: "Coming Through",
		Text: "Knockback distance of Boosters is increased by 100%.",
	},
	"DVaBoostersCrashCourse": {
		Name: "Crash Course",
		Text: "Quest: Damage enemy Heroes with Boosters.\nReward: After damaging 20 Heroes, each use of Boosters lowers its cooldown by 1 second per Hero hit.",
	},
	"DVaBoostersHitTheNitrous": {
		Name: "Hit the Nitrous",
		Text: "Initial speed bonus of Boosters is increased to 325%, then decays to normal boost speed over 0.5 seconds. During this time, Boosters deals 50% extra damage.",
	},
	"DVaBoostersRushdown": {
		Name: "Rush-down",
		Text: "If D.Va neither takes nor deals damage during Boosters, its cooldown is lowered by 7 seconds.",
	},
	"DVaBunnyHop": {
		Name: "Bunny Hop",
		Text: "D.Va's Mech becomes Unstoppable and stomps every 0.5 seconds, dealing 60 damage and Slowing enemies by 40%. Lasts 4 seconds.\nRequires Mech Mode.",
	},
	"DVaBunnyHopStopAndPop": {
		Name: "Stop and Pop",
		Text: "Bunny Hop deals 200% more damage if D.Va isn't moving at the moment of impact.",
	},
	"DVaDefenseMatrixAggressionMatrix": {
		Name: "Aggression Matrix",
		Text: "Basic Attacks in Mech Mode against Heroes lower the cooldown of Defense Matrix by 0.2 seconds.",
	},
	"DVaDefenseMatrixDazerZone": {
		Name: "Dazer Zone",
		Text: "Enemy Heroes affected by Defense Matrix are Slowed by 20%.",
	},
	"DVaDefenseMatrixDivertingPower": {
		Name: "Diverting Power",
		Text: "The area of Defense Matrix is 100% wider and 25% longer, but D.Va's Movement Speed is reduced by 30% for the duration.",
	},
	"DVaDefenseMatrixFusionGenerator": {
		Name: "Fusion Generator",
		Text: "Every time an enemy Hero deals damage while inside Defense Matrix, D.Va's Self-Destruct Charge increases by 1%.\nMax of 15% Charge gained per use.",
	},
	"DVaDefenseMatrixGetThroughThis": {
		Name: "Get Through This!",
		Text: "Increase the duration of Defense Matrix by 3 seconds.",
	},
	"DVaMechAblativeArmor": {
		Name: "Ablative Armor",
		Text: "Damage against D.Va's Mech that would deal 4.1% or less of its maximum Health is reduced by 50%.\nThis does not decrease the amount of Self-Destruct Charge gained.",
	},
	"DVaMechEmergencyShielding": {
		Name: "Emergency Shielding",
		Text: "When D.Va's Mech would be destroyed, it instead gains a Shield that absorbs 209.5 damage over 6 seconds.\nThis effect has a 15 second cooldown.",
	},
	"DVaMechExpensivePlating": {
		Name: "Expensive Plating",
		Text: "Increase Mech Health by 20%, but also increase the cooldown of Call Mech by 15 seconds.\nThis does not decrease the amount of Self-Destruct Charge gained.",
	},
	"DVaMechProMoves": {
		Name: "Pro Moves",
		Text: "D.Va's Mech gains 2% Movement Speed for 1.5 seconds every time it takes damage, up to 30%.",
	},
	"DVaPilotCallMechMEKAfall": {
		Name: "MEKAfall",
		Text: "Call Mech is instant and can target a location. Upon landing, the Mech deals 202 damage to enemies in the impact area.",
	},
	"DVaPilotConcussivePulse": {
		Name: "Concussive Pulse",
		Text: "D.Va gains the Concussive Pulse ability, allowing her to deal 141 damage to enemies in a cone and knock them back.\nRequires Pilot Mode.",
	},
	"DVaPilotGGWP": {
		Name: "GG, WP",
		Text: "Increase Pilot Mode Basic Attack damage by 50%. Participating in a Takedown while in Pilot Mode instantly refreshes the cooldown of Call Mech.",
	},
	"DVaPilotNanoweaveSuit": {
		Name: "Nanoweave Suit",
		Text: "For 4 seconds after ejecting from her Mech, D.Va gains 50 Armor and her Basic Attacks grant 50% more cooldown reduction towards Call Mech.",
	},
	"DVaPilotSuppressingFire": {
		Name: "Suppressing Fire",
		Text: "Pilot Mode Basic Attacks have 20% increased range and Slow enemy Movement Speed by 20% for 2.5 seconds.",
	},
	"DVaPilotTorpedoDash": {
		Name: "Torpedo Dash",
		Text: "D.Va gains the Torpedo Dash ability, allowing her to dash forward and pass through enemies.\nRequires Pilot Mode.",
	},
	"DVaSelfDestructBringItOn": {
		Name: "Bring It On",
		Text: "Self-Destruct's Charge amount gained from losing Mech Health is increased by 30%.",
	},
	"DVaSelfDestructNuclearOption": {
		Name: "Nuclear Option",
		Text: "Increase Self Destruct's detonation timer by 3 seconds. Detonation damage increased by 70%.\nA new Mech can still be called after 4 seconds.",
	},
	"DehakaAdaptationChangeIsSurvivalTalent": {
		Name: "Change Is Survival",
		Text: "Increases Adaptation healing to 200% of the damage received and reduces the cooldown by 30 seconds.",
	},
	"DehakaBrushstalkerApexPredator": {
		Name: "Apex Predator",
		Text: "Reduces Brushstalker's cooldown by 25 seconds and the cast time by 0.5 seconds.",
	},
	"DehakaBrushstalkerFerociousStalker": {
		Name: "Ferocious Stalker",
		Text: "Dark Swarm deals 40% more damage while Brushstalker's Movement Speed bonus is active.",
	},
	"DehakaBurrowTalentLurkerStrain": {
		Name: "Lurker Strain",
		Text: "Emerging from Burrow grants Dehaka Stealth for 3 seconds and also knocks nearby enemies back, Slowing them by 30% for 2 seconds.",
	},
	"DehakaDarkSwarmTalentEnduringSwarm": {
		Name: "Enduring Swarm",
		Text: "Dark Swarm grants 40 Spell Armor while active.",
	},
	"DehakaDarkSwarmTalentPrimalAggression": {
		Name: "Primal Aggression",
		Text: "Dark Swarm duration increased by .5 seconds and it deals 100% bonus damage to Minions and Mercenaries. ",
	},
	"DehakaDarkSwarmTalentSymbiosis": {
		Name: "Symbiosis",
		Text: "Every time Dark Swarm damages an enemy Hero, reduce its cooldown by 0.4 seconds.",
	},
	"DehakaDragTalentConstriction": {
		Name: "Constriction",
		Text: "Increases Drag duration by 0.5 seconds.",
	},
	"DehakaDragTalentElongatedTongue": {
		Name: "Elongated Tongue",
		Text: "Increase Drag range by 20%.",
	},
	"DehakaDragTalentFeedingFrenzy": {
		Name: "Feeding Frenzy",
		Text: "Basic Attacks reduce Drag's cooldown by 2 seconds.",
	},
	"DehakaDragTalentParalyzingEnzymes": {
		Name: "Paralyzing Enzymes",
		Text: "Drag slows enemies by 50% for 2 seconds after it ends.",
	},
	"DehakaEssenceClaws": {
		Name: "Essence Claws",
		Text: "Dehaka's Basic Attacks slow the target by 20% for 1 second. If the target is a Hero, Dehaka gains 5 Essence.",
	},
	"DehakaEssenceCollectionOneWhoCollectsTalent": {
		Name: "One-Who-Collects",
		Text: "Increases Essence collected from Minions by 50%.",
	},
	"DehakaEssenceCollectionTalentEnhancedAgility": {
		Name: "Enhanced Agility",
		Text: "Brushstalker's Movement Speed bonus now lasts for 5 seconds after leaving a bush.\nRepeatable Quest: Every 50 Essence collected permanently increases the Movement Speed bonus of Brushstalker by 2%, up to 20%.",
	},
	"DehakaEssenceCollectionTalentEssenceDevourer": {
		Name: "Essence Devourer",
		Text: "Quest: Regeneration Globes grant 11 Essence and increases maximum Essence by 1, up to 10.\nReward: After collecting 10, Regeneration Globes will instead grant 16 Essence.",
	},
	"DehakaEssenceCollectionTalentHeroStalker": {
		Name: "Hero Stalker",
		Text: "Increases Essence gained from Takedowns by 100%.",
	},
	"DehakaEssenceCollectionTalentTissueRegeneration": {
		Name: "Tissue Regeneration",
		Text: "Regeneration Globes grant 10 Essence.\nRepeatable Quest: Every 50 Essence collected permanently increases Health Regeneration by 4, up to 40, and maximum Essence by 1, up to 10.",
	},
	"DehakaHeroicAbilityAdaptation": {
		Name: "Adaptation",
		Text: "After 4 seconds, heal for 100% of the damage Dehaka took over this period.",
	},
	"DehakaHeroicAbilityIsolation": {
		Name: "Isolation",
		Text: "Launch biomass that hits the first enemy Hero dealing 200 damage, Silencing and Slowing them 30% for 3 seconds. Additionally, their vision radius is greatly reduced for 6 seconds.",
	},
	"DehakaIsolationTalentContagion": {
		Name: "Contagion",
		Text: "Isolation hits all Heroes near the first target.",
	},
	"DehakaPrimalRage": {
		Name: "Primal Rage",
		Text: "Gain 1% increased Attack Damage per Essence stored.",
	},
	"DehakaPrimalSwarm": {
		Name: "Primal Swarm",
		Text: "Dark Swarm causes enemies hit to lose 10 Armor for 0.8 seconds, causing them to take 10% extra damage.",
	},
	"DehakaSwiftPursuit": {
		Name: "Swift Pursuit",
		Text: "Brushstalker's bush movement speed bonus now lasts for 4 seconds after leaving the bush.",
	},
	"DehakaTunnelingClaws": {
		Name: "Tunneling Claws",
		Text: "Dehaka can move while Burrowed.",
	},
	"DemonHunterCombatStyleHotPursuit": {
		Name: "Hot Pursuit",
		Text: "When at 10 stacks of Hatred, the Movement Speed bonus increases to 20% total and Valla gains 4 Mana per second.",
	},
	"DemonHunterCombatStyleRancor": {
		Name: "Rancor",
		Text: "Each stack of Hatred also increases Attack Speed by 2%. While at 10 stacks of Hatred, gain an additional 30% Attack Speed.",
	},
	"DemonHunterCombatStyleTemperedByDiscipline": {
		Name: "Tempered by Discipline",
		Text: "While at 10 stacks of Hatred, Basic Attacks heal for 25% of the damage dealt.",
	},
	"DemonHunterCreedoftheHunter": {
		Name: "Creed of the Hunter",
		Text: "Increases Attack Speed by 10%.\nQuest: Use 100 Basic Attacks against Heroes.\nReward: Hatred grants an additional 2% Basic Attack Damage per stack.",
	},
	"DemonHunterDeathDealer": {
		Name: "Death Dealer",
		Text: "Increases Vault Basic Attack damage bonus from 6% to 14% per stack of Hatred. If this attack kills its victim, the Mana cost and cooldown of Vault are refunded.",
	},
	"DemonHunterFarflightQuiver": {
		Name: "Farflight Quiver",
		Text: "Increases Valla's Basic Attack range by 2.2.",
	},
	"DemonHunterGloom": {
		Name: "Gloom",
		Text: "Activate to consume all Hatred, granting 3 Spell Armor per Hatred consumed for 5 seconds.\nPassive: Permanently gain 20 Spell Armor, reducing Ability Damage taken by 20%.",
	},
	"DemonHunterHeroicAbilityRainofVengeance": {
		Name: "Rain of Vengeance",
		Text: "Launch a wave of Shadow Beasts that deals 250 damage and stuns enemies in the target area for 0.5 seconds.\nStores up to 2 charges.",
	},
	"DemonHunterHeroicAbilityStrafe": {
		Name: "Strafe",
		Text: "Rapidly attack enemies within 10 range for 120 damage per hit, prioritizing Heroes over Minions. Valla is able to move and use Vault while strafing. Lasts for 4 seconds.",
	},
	"DemonHunterManticore": {
		Name: "Manticore",
		Text: "Every 3rd Basic Attack against the same Heroic target deals an additional 5% of their maximum Health as damage.",
	},
	"DemonHunterMasteryArsenal": {
		Name: "Arsenal",
		Text: "Multishot also fires 3 grenades which deal 120 damage. Each enemy Hero hit by a grenade reduces the cooldown of Multishot by 2 seconds.",
	},
	"DemonHunterMasteryCaltrops": {
		Name: "Caltrops",
		Text: "Drop 3 Caltrops while Vaulting. Caltrops do 60 damage and slow enemies by 25% for 1 second.\nQuest: Spend 160 seconds at 10 stacks of Hatred.\nReward: Vault cooldown is reduced by 5 seconds.",
	},
	"DemonHunterMasteryCompositeArrowsMultishot": {
		Name: "Composite Arrows",
		Text: "Increases the range of Multishot by 20%.",
	},
	"DemonHunterMasteryDeathSiphon": {
		Name: "Death Siphon",
		Text: "Strafe also fires penetrating bolts in a line for 34 damage every 0.25 seconds. 25% of the damage dealt with Strafe is returned as Health.",
	},
	"DemonHunterMasteryFrostShot": {
		Name: "Frost Shot",
		Text: "Increases the range of Multishot by 20% and it also slows by 25% for 1.5 seconds.",
	},
	"DemonHunterMasteryMonsterHunterHungeringArrow": {
		Name: "Monster Hunter",
		Text: "Reduces the Mana cost of Hungering Arrow by 40 and increases the damage it deals to Minions, Mercenaries, and Monsters by 150%.",
	},
	"DemonHunterMasteryPuncturingArrow": {
		Name: "Puncturing Arrow",
		Text: "Quest: Hitting an enemy Hero with the initial impact of Hungering Arrow increases its damage by 5, up to a maximum of 100.\nReward: After gaining 100 bonus damage, gain an additional 75 bonus damage and Hungering Arrow also bounces 1 additional time.",
	},
	"DemonHunterMasteryRepeatingArrowVault": {
		Name: "Repeating Arrow",
		Text: "The cooldown for Hungering Arrow is reset when Vault is used.",
	},
	"DemonHunterMasterySiphoningArrow": {
		Name: "Siphoning Arrow",
		Text: "Valla heals for 75% of the damage dealt to Heroes by Hungering Arrow.",
	},
	"DemonHunterMasteryStormofVengeance": {
		Name: "Storm of Vengeance",
		Text: "Basic Attacks reduce the charge cooldown of Rain of Vengeance by 5 seconds.",
	},
	"DemonHunterMasteryTumble": {
		Name: "Tumble",
		Text: "Vault gains an additional charge, allowing it to be cast twice in quick succession.",
	},
	"DemonHunterPunishment": {
		Name: "Punishment",
		Text: "Each stack of Hatred increases the damage of Multishot by 4%.",
	},
	"DemonHunterSeethingHatred": {
		Name: "Seething Hatred",
		Text: "Basic Attacks now grant 2 stacks of Hatred and the time it takes before Hatred expires is increased by 2 seconds.",
	},
	"DiabloBulwark": {
		Name: "Bulwark",
		Text: "Increases the duration of Armor granted by Shadow Charge from 2 seconds to 4 seconds.",
	},
	"DiabloDebilitatingFlames": {
		Name: "Debilitating Flames",
		Text: "Enemy Heroes hit by Fire Stomp are slowed by 15% for 2 seconds, up to 30%.",
	},
	"DiabloDevastatingCharge": {
		Name: "Devastating Charge",
		Text: "Shadow Charge's terrain collision against Heroes deals an additional 5% of the target's maximum life.\nQuest: Each time an enemy Hero takes terrain collision damage, increase the maximum life damage by 2%, to a maximum of 10% bonus.",
	},
	"DiabloDiabolicalMomentum": {
		Name: "Diabolical Momentum",
		Text: "Basic Attacks reduce the cooldown of Overpower and Shadow Charge by 1.5 seconds.",
	},
	"DiabloFearfulPresence": {
		Name: "Fearful Presence",
		Text: "Activate to reduce enemy damage",
	},
	"DiabloHellfire": {
		Name: "Hellfire",
		Text: "Each enemy Hero hit by Fire Stomp increases the damage of the next Fire Stomp by 12%, up to 120%.",
	},
	"DiabloHellgate": {
		Name: "Hellgate",
		Text: "Teleport and place a demonic rune at target location. After 1.8 seconds the rune explodes dealing 137 damage and stunning enemies for 1.8 seconds.",
	},
	"DiabloHeroicAbilityApocalypse": {
		Name: "Apocalypse",
		Text: "Create a demonic rune under each enemy Hero on the battleground. After 1.8 seconds the rune explodes dealing 137 damage and stunning them for 1.8 seconds.",
	},
	"DiabloHeroicAbilityLightningBreath": {
		Name: "Lightning Breath",
		Text: "Become Unstoppable while channeling lightning that deals 800 damage over 4 seconds.  The direction of the Lightning changes with your mouse cursor position.",
	},
	"DiabloLifeLeech": {
		Name: "Life Leech",
		Text: "Basic Attacks against enemy Heroes deal bonus damage equal to 1% of the Hero's maximum Health and heal Diablo for the same amount.",
	},
	"DiabloMasteryDemonicStrength": {
		Name: "Demonic Strength",
		Text: "Once Overpower's stun expires, the target is slowed by 30% for 3 seconds.",
	},
	"DiabloMasteryDevilsDueBlackSoulstone": {
		Name: "Devil's Due",
		Text: "Black Soulstone increases the effects of Regen Globes and Healing Fountains by 2% per Soul.",
	},
	"DiabloMasteryDyingBreathApocalypse": {
		Name: "Dying Breath",
		Text: "Apocalypse's cooldown is reduced by 20 seconds and is cast for free when Diablo dies.",
	},
	"DiabloMasteryFireDevilFireStomp": {
		Name: "Fire Devil",
		Text: "Fire Stomp increases Diablo's Basic Attack damage by 20% and surrounds Diablo in flames that deal 20 damage every second. Lasts 6 seconds.",
	},
	"DiabloMasteryFromTheShadowsShadowCharge": {
		Name: "From the Shadows",
		Text: "Increases the cast range of Shadow Charge by 40%.",
	},
	"DiabloMasteryHellstormLightningBreath": {
		Name: "Hellstorm",
		Text: "Lightning Breath lasts and reaches 50% longer.",
	},
	"DiabloMasterySoulFeastBlackSoulstone": {
		Name: "Soul Feast",
		Text: "Black Soulstone increases Diablo's Health Regeneration by 0.4 per second per Soul.",
	},
	"DiabloSoulShield": {
		Name: "Soul Shield",
		Text: "Black Soulstone grants Diablo 0.25 Spell Armor per soul, reducing Ability Damage taken by 0.25% per Soul.",
	},
	"DiabloSpeedDemon": {
		Name: "Speed Demon",
		Text: "Being Stunned or Rooted increases Diablo's Movement Speed by 30% for 5 seconds.",
	},
	"DiabloTalentDominationOverpower": {
		Name: "Domination",
		Text: "Casting Overpower resets the cooldown of Shadow Charge.",
	},
	"DiabloTalentLordOfTerror": {
		Name: "Lord of Terror",
		Text: "Activate to steal 10% of the maximum Health of nearby enemy Heroes.",
	},
	"DryadAbolishMagic": {
		Name: "Abolish Magic",
		Text: "Target an Ally to remove all damage over time and disabling effects from them and Lunara. For 2 seconds after, the duration of disabling effects is reduced by 50%.",
	},
	"DryadBlossomSwell": {
		Name: "Blossom Swell",
		Text: "Increases Noxious Blossom's radius by 20%.",
	},
	"DryadBoundlessStrideTalent": {
		Name: "Boundless Stride",
		Text: "All Leaping Strike charges are returned every 20 seconds. Leaping Strike can be used on allies.",
	},
	"DryadChokingPollen": {
		Name: "Choking Pollen",
		Text: "Increases Noxious Blossom damage by 125% against enemies afflicted with Nature's Toxin.",
	},
	"DryadCruelSpores": {
		Name: "Cruel Spores",
		Text: "When Crippling Spores hits only Minions or Mercenaries, its cooldown is reduced by 80% and 20 Mana is restored.",
	},
	"DryadDividingWisp": {
		Name: "Dividing Wisp",
		Text: "The first time a Wisp is relocated, it leaves a copy of itself in the previous location.",
	},
	"DryadForestsWrath": {
		Name: "Forest's Wrath",
		Text: "Increases Thornwood Vine's range by 30%, its speed by 20%, and Lunara's vision radius by 35%.",
	},
	"DryadGallopingGait": {
		Name: "Galloping Gait",
		Text: "Activate to increase the Movement Speed bonus of Dryad's Swiftness to 80% for 6 seconds.",
	},
	"DryadGreaterSpellShield": {
		Name: "Greater Spell Shield",
		Text: "Every 30 seconds, gain 75 Spell Armor against the next enemy Ability and subsequent Abilities for 3 seconds, reducing the damage taken by 75%.\nCan be toggled to allow or prevent this talent from triggering automatically.",
	},
	"DryadHeroicAbilityLeapingStrike": {
		Name: "Leaping Strike",
		Text: "Leap over an enemy, slowing them by 80% for 0.3 seconds and dealing 271 damage.\nStores up to 2 charges.",
	},
	"DryadHeroicAbilityThornwoodVine": {
		Name: "Thornwood Vine",
		Text: "Send forth vines that deal 164 damage to all enemies in a line.\nStores up to 3 charges.",
	},
	"DryadInvigoratingSpores": {
		Name: "Invigorating Spores",
		Text: "Using Crippling Spores increases Lunara's Attack Speed by 50% for 6 seconds.",
	},
	"DryadLetThemWither": {
		Name: "Let Them Wither",
		Text: "Increases Crippling Spores' slow to 50% and causes it to no longer decay.",
	},
	"DryadNaturalPerspective": {
		Name: "Natural Perspective",
		Text: "Nature's Toxin reveals enemies for its duration.",
	},
	"DryadNaturesCulling": {
		Name: "Nature's Culling",
		Text: "Increases Nature's Toxin's damage by 150% to non-Heroes.",
	},
	"DryadNimbleWisp": {
		Name: "Nimble Wisp",
		Text: "Increases Wisp Movement Speed by 100% and vision radius by 50%.",
	},
	"DryadPesteringBlossom": {
		Name: "Pestering Blossom",
		Text: "Increases Noxious Blossom's range by 50%.",
	},
	"DryadPhotosynthesis": {
		Name: "Photosynthesis",
		Text: "Crippling Spores returns 10 Mana per enemy hit, up to 40 Mana.",
	},
	"DryadSiphoningToxin": {
		Name: "Siphoning Toxin",
		Text: "As long as Nature's Toxin is active on an enemy, heal 18 Health a second.",
	},
	"DryadSkyboundWisp": {
		Name: "Skybound Wisp",
		Text: "Wisp can see over obstacles and reveals the area for 4 seconds after dying.",
	},
	"DryadSplinteredSpear": {
		Name: "Splintered Spear",
		Text: "Using Noxious Blossom causes Lunara's next Basic Attack to hit up to 4 enemies. These extra attacks can apply Nature's Toxin.",
	},
	"DryadStarWoodSpear": {
		Name: "Star Wood Spear",
		Text: "Using Crippling Spores increases Lunara's Basic Attack range by 2.8 for 6 seconds.",
	},
	"DryadTimelostWisp": {
		Name: "Timelost Wisp",
		Text: "Wisp costs no Mana and its cooldown recharges 125% faster while no Wisp is active.",
	},
	"DryadUnfairAdvantage": {
		Name: "Unfair Advantage",
		Text: "Nature's Toxin deals 50% more damage to Heroes that are slowed.",
	},
	"DryadWildVigor": {
		Name: "Wild Vigor",
		Text: "Using Crippling Spores increases the damage of Lunara's next 4 Basic Attacks by 50%.",
	},
	"ETCBlockParty": {
		Name: "Block Party",
		Text: "Using a Basic or Heroic ability also gives nearby allied Heroes a stack of Block, granting 50 Physical Armor against the next Hero Basic Attack for 8 seconds, reducing the damage taken by 50%. Maximum 2 stacks of Block.",
	},
	"ETCCombatStyleEchoPedal": {
		Name: "Echo Pedal",
		Text: "Using a Basic or Heroic Ability releases two pulses of 15 damage.  This deals 250% bonus damage to Minions and Mercenaries. The first occurs instantly, the second occurs 2 seconds later.",
	},
	"ETCCrowdSurfer": {
		Name: "Crowd Surfer",
		Text: "Allows Powerslide to travel over walls and terrain.",
	},
	"ETCHeroicAbilityMoshPit": {
		Name: "Mosh Pit",
		Text: "After 0.8 seconds, channel to stun nearby enemies for 4 seconds.",
	},
	"ETCHeroicAbilityStageDive": {
		Name: "Stage Dive",
		Text: "Leap to target location, landing after 2.75 seconds, dealing 660 damage to enemies in the area, and slowing them by 50% for 3 seconds.",
	},
	"ETCMasteryEncore": {
		Name: "Encore",
		Text: "Face Melt leaves an Amp behind, which will knock enemies away again 2 seconds later.",
	},
	"ETCMasteryFaceSmelt": {
		Name: "Face Smelt",
		Text: "Face Melt slows enemies by 49.9% fading over 2 seconds.",
	},
	"ETCMasteryGuitarHero": {
		Name: "Guitar Hero",
		Text: "While Guitar Solo is active, E.T.C. heals for 50% of his damage dealt by Basic Attacks.",
	},
	"ETCMasteryGuitarSoloAggressiveShredding": {
		Name: "Aggressive Shredding",
		Text: "E.T.C.'s Basic Attacks reduce the cooldown of Guitar Solo by 0.8 seconds.",
	},
	"ETCMasteryHammeron": {
		Name: "Hammer-on",
		Text: "After using an Ability, E.T.C.'s next 2 Basic Attacks deal 35% more damage.",
	},
	"ETCMasteryJustKeepRocking": {
		Name: "Just Keep Rockin'",
		Text: "While Guitar Solo is active, the duration of stuns, slows, roots, and polymorphs are reduced by 50%.",
	},
	"ETCMasteryMicCheck": {
		Name: "Mic Check",
		Text: "Hitting at least 2 targets with Face Melt reduces its cooldown by 6 seconds.",
	},
	"ETCMasteryProgRock": {
		Name: "Prog Rock",
		Text: "Quest: Gathering a Regeneration Globe permanently increases the healing per second of Guitar Solo by 5, to a maximum of 100\nReward: After gathering 20 Regeneration Globes, Guitar Solo also heals nearby allied Heroes for 50 Health every second while it is active.",
	},
	"ETCMasteryRollingLikeaStone": {
		Name: "Rolling Like a Stone",
		Text: "Increases the range of Powerslide by 25%.",
	},
	"ETCMasteryShowStopper": {
		Name: "Show Stopper",
		Text: "After using Powerslide, gain 15 Armor for 4 seconds, reducing all damage taken by 15%.",
	},
	"ETCMasterySpeedMetal": {
		Name: "Speed Metal",
		Text: "Using a Basic or Heroic ability also gives nearby allied Heroes 15% Movement Speed for 2 seconds.",
	},
	"FaerieDragonHardenedFocus": {
		Name: "Hardened Focus",
		Text: "While above 80% life, your Basic Ability cooldowns regenerate 50% faster.",
	},
	"FaerieDragonHeroicAbilityBlinkHeal": {
		Name: "Blink Heal",
		Text: "Teleport to a nearby ally, healing them for 204.\nStores up to 2 charges.",
	},
	"FaerieDragonHeroicAbilityEmeraldWind": {
		Name: "Emerald Wind",
		Text: "After 0.5 seconds, create an expanding nova of wind, dealing 319 damage and pushing enemies away. \nPassive: Increases the healing of Soothing Mist by 5%.",
	},
	"FaerieDragonMasteryArcanePrecision": {
		Name: "Arcane Precision",
		Text: "Increases Arcane Flare's inner area damage by 50%.",
	},
	"FaerieDragonMasteryContinuousWinds": {
		Name: "Continuous Winds",
		Text: "Emerald Wind releases two additional novas that deal 25% damage. Also further increases the healing bonus of Soothing Mist by 7%.",
	},
	"FaerieDragonMasteryCritterize": {
		Name: "Critterize",
		Text: "While Polymorphed, the target's Armor is reduced by 25, increasing the damage they take by 25%.",
	},
	"FaerieDragonMasteryPhaseShield": {
		Name: "Phase Shield",
		Text: "After finishing the teleport, both Brightwing and her target gain a 350 point Shield for 10 seconds.",
	},
	"FaerieDragonMasteryShieldDust": {
		Name: "Shield Dust",
		Text: "Pixie Dust grants 50 Physical Armor for 3 seconds, reducing Physical damage taken by 50%.",
	},
	"FaerieDragonMasteryStickyFlare": {
		Name: "Sticky Flare",
		Text: "Enemies hit by Arcane Flare have their Movement Speed slowed by 40% for 2 seconds.",
	},
	"FalstadHammerangBOOMerang": {
		Name: "BOOMerang",
		Text: "Reactivate Hammerang mid-flight to deal 165 damage around the hammer.",
	},
	"FalstadHammerangGatheringStorm": {
		Name: "Gathering Storm",
		Text: "Quest: Every time Hammerang hits a Hero increase its damage by 1.5.",
	},
	"FalstadHeroicAbilityHinterlandBlast": {
		Name: "Hinterland Blast",
		Text: "After 1 second, deal 475 damage to enemies within a long line. The cooldown is reduced by 25 seconds for every enemy Hero hit.",
	},
	"FalstadHeroicAbilityMightyGust": {
		Name: "Mighty Gust",
		Text: "Push enemies away, and slow their Movement Speed by 40% decaying over 4 seconds.",
	},
	"FalstadMasteryAerieGustsTailwind": {
		Name: "Aerie Gusts",
		Text: "Reduces the activation time for Tailwind from 6 to 3 seconds, and increases the Movement Speed bonus from 15% to 25%.",
	},
	"FalstadMasteryAfterburner": {
		Name: "Afterburner",
		Text: "Barrel Roll increases Movement Speed by 60%, decaying over 3 seconds.",
	},
	"FalstadMasteryBarrelRollFlowRider": {
		Name: "Flow Rider",
		Text: "While Tailwind is active, Falstad's Basic Abilities recharge 100% faster.",
	},
	"FalstadMasteryBarrelRollFreeRoll": {
		Name: "Free Roll",
		Text: "Barrel Roll costs no Mana.",
	},
	"FalstadMasteryCalloftheWildhammerHinterlandBlast": {
		Name: "Call of the Wildhammer",
		Text: "Hinterland Blast has double the range and deals 25% more damage.",
	},
	"FalstadMasteryCripplingHammerHammerang": {
		Name: "Crippling Hammer",
		Text: "Increases Hammerang's slow from 25% to 50%.",
	},
	"FalstadMasteryDogFightBarrelRoll": {
		Name: "Dog Fight",
		Text: "Dog Fight \nThe Barrel Roll Shield lasts 3 seconds longer.",
	},
	"FalstadMasteryFlightEpicMount": {
		Name: "Epic Mount",
		Text: "Reduce the cooldown of Flight to 30 seconds, reduce the cast time before flying to 0.5 seconds, and increase the speed by 50%.",
	},
	"FalstadMasteryFlyAwayFlight": {
		Name: "Fly Away!",
		Text: "Fly Away! \nThe cooldown for Flight is reduced by 15 seconds.",
	},
	"FalstadMasteryHammerangPowerThrow": {
		Name: "Power Throw",
		Text: "Increase the range and slow duration of Hammerang by 40%.",
	},
	"FalstadMasteryHammerangSecretWeapon": {
		Name: "Secret Weapon",
		Text: "Increases Hammerang's range by 30% and Basic Attacks deal 60% bonus damage while Hammerang is in flight.",
	},
	"FalstadMasteryLightningRodChargedUp": {
		Name: "Charged Up",
		Text: "Increases the number of Lightning Rod strikes by 2 and its range by 25%. ",
	},
	"FalstadMasteryLightningRodLightningChain": {
		Name: "Lightning Chain",
		Text: "Lightning Chain \nLightning Rod chains to up to 2 additional targets for 25% damage each. ",
	},
	"FalstadMasteryLightningRodStaticShield": {
		Name: "Static Shield",
		Text: "Gain a Shield equal to 4% of Falstad's maximum Health after every Lightning Rod strike. Lasts 4 seconds and stacks.",
	},
	"FalstadMasteryLightningRodThunderstrikes": {
		Name: "Thunderstrikes",
		Text: "Lightning Rod deals 20% more damage each subsequent strike.",
	},
	"FalstadMasteryMightyGustWindTunnel": {
		Name: "Wind Tunnel",
		Text: "Mighty Gust creates a wind tunnel for 4 seconds. Enemies caught in the tunnel will periodically be pushed back.",
	},
	"FalstadMasteryUpdraftBarrelRoll": {
		Name: "Updraft",
		Text: "Increases Barrel Roll's range, Shield amount, and Shield duration by 40%.",
	},
	"FalstadTalentHammerGains": {
		Name: "Hammer Gains",
		Text: "Basic Attacks heal for 0% of the damage dealt to the primary target.",
	},
	"FalstadWingman": {
		Name: "Wingman",
		Text: "Enemy Minions killed near Falstad grant a stack of Bribe. Falstad can use 20 stacks to bribe a Mercenary, instantly defeating them and permanently increasing the damage of Lightning Rod by 5%. Does not work on Bosses. Maximum of 80 stacks.",
	},
	"GallBombsAway": {
		Name: "Bomb's Away",
		Text: "Quest: Damaging a Hero with Runic Blast increases its damage by 5.\nReward: After damaging 20 Heroes, the range and speed of Cho's Rune Bomb is increased by 20%.",
	},
	"GallDeafeningBlast": {
		Name: "Deafening Blast",
		Text: "Runic Blast silences enemy Heroes for 1.2 seconds.",
	},
	"GallDoubleBack": {
		Name: "Double Back",
		Text: "Activate Dread Orb before the second bounce ends to reverse the direction of the third bounce.\nPassive: The third bounce of Dread Orb deals 200% more damage to Non-Heroic targets. ",
	},
	"GallDoubleTrouble": {
		Name: "Double Trouble",
		Text: "Quest: If Shadowflame hits an enemy Hero that is afflicted by Cho's Consuming Blaze, its cooldown is reduced by 1 second.\nReward: After hitting 20 Heroes, Shadowflame's cooldown is instead permanently reduced by 1 second.",
	},
	"GallEdgeOfMadness": {
		Name: "Edge of Madness",
		Text: "Every subsequent hit of Shadowflame against the same enemy Hero deals an additional 8% damage, to a maximum of 40%. These bonuses are lost if the Hero has not been hit for 4 seconds.",
	},
	"GallEyeOfKilrogg": {
		Name: "Eye of Kilrogg",
		Text: "Place an eye, granting vision of a large area around it for 45 seconds. The eye can be killed by enemies with 2 Basic Attacks. Stores up to 2 charges.",
	},
	"GallHeroicAbilityShadowboltVolley": {
		Name: "Shadow Bolt Volley",
		Text: "After 1 second, unleash 19.4 Shadow Bolts over 3.9 seconds, each dealing 87 damage to the first target hit. The bolts fire towards your mouse.",
	},
	"GallHeroicAbilityTwistingNether": {
		Name: "Twisting Nether",
		Text: "After 1 second, nearby enemies are slowed by 50% while Gall channels, up to 5 seconds. Activate to deal 353 damage.",
	},
	"GallKeepMoving": {
		Name: "Keep Moving!",
		Text: "Increases the duration of Shove's movement speed bonus by 1 second. Cho's Basic Attacks against Heroes increases the duration of an active Shove by 1 second.",
	},
	"GallLeadenOrb": {
		Name: "Leaden Orb",
		Text: "Dread Orb stuns enemy Heroes for 0.5 seconds.",
	},
	"GallOgreRampage": {
		Name: "Ogre Rampage",
		Text: "Gall's Basic Abilities cooldown 75% faster for 5 seconds after Cho activates Ogre Hide.",
	},
	"GallPsychoticBreak": {
		Name: "Psychotic Break",
		Text: "Upon dying, Gall gains Ogre Rage and can use Abilities for 10 seconds.",
	},
	"GallRisingDread": {
		Name: "Rising Dread",
		Text: "Each bounce of Dread Orb increases its radius by 25% and damage by 25%.",
	},
	"GallRunicPersistence": {
		Name: "Runic Persistence",
		Text: "Runic Blast deals an additional 180 damage over 3 seconds in an area around the detonation point.",
	},
	"GallSearingShadows": {
		Name: "Searing Shadows",
		Text: "Enemy Heroes hit by Shadowflame take an additional 2% of their maximum Health as damage.",
	},
	"GallShadowboltVolleyShadowfury": {
		Name: "Shadowfury",
		Text: "Shadow Bolt Volley hits all enemies in its path.",
	},
	"GallShadowsnare": {
		Name: "Shadowsnare",
		Text: "Shadowflame slows enemy Heroes by 10% for 4 seconds. This effect can stack up to 3 times.",
	},
	"GallShiftingNether": {
		Name: "Shifting Nether",
		Text: "Twisting Nether now teleports Cho to a targeted location before it begins channeling.",
	},
	"GallShove": {
		Name: "Shove",
		Text: "Nudge Cho a small distance and grant him 25% Movement Speed for 2 seconds.",
	},
	"GallTaskmaster": {
		Name: "Taskmaster",
		Text: "Reduces the cooldown of Shove by 10 seconds. Every time Cho is hit by a Hero Basic Attack, reduce the Cooldown of Shove by 1 second.",
	},
	"GallTheWillofGall": {
		Name: "The Will of Gall",
		Text: "Takedowns permanently increase the bonus of Ogre Rage by 2%.",
	},
	"GallTwilightNova": {
		Name: "Twilight Nova",
		Text: "After the first bounce of Dread Orb, 2 extra bombs bounce to the sides once.",
	},
	"GallWeSeeYou": {
		Name: "We See You!",
		Text: "Eye of Kilrogg's cooldown is decreased by 15 seconds and its range is increased by 33.3%.",
	},
	"GarroshArmorUpBodyCheck": {
		Name: "Body Check",
		Text: "Activate to deal 111 damage to a target enemy and Slow them by 30% for 2 seconds.\nDamage is increased by 200% of Armor Up's current bonus, and the Slow amount is increased by 30% if Armor Up's bonus is above 25.",
	},
	"GarroshArmorUpDoubleUp": {
		Name: "Double Up",
		Text: "Armor Up can be activated to increase its Armor bonus by 0% for 3 seconds.",
	},
	"GarroshArmorUpInnerRage": {
		Name: "Inner Rage",
		Text: "Body Check gains an additional charge, and its cooldown recharges 200% faster while Armor Up's bonus is at or above 25.",
	},
	"GarroshBloodthirstBloodcraze": {
		Name: "Bloodcraze",
		Text: "When hitting a Hero, Bloodthirst heals for an additional 9.9% of Garrosh's maximum Health over 3 seconds.",
	},
	"GarroshBloodthirstInFortheKill": {
		Name: "In For the Kill",
		Text: "Increase Bloodthirst's damage against non-Heroes by 70%. Killing enemies with Bloodthirst resets its cooldown and refunds its Mana cost.",
	},
	"GarroshBloodthirstThirstforBattle": {
		Name: "Thirst for Battle",
		Text: "Basic Attacks against Heroes reduce the cooldown of Bloodthirst by 2 seconds.",
	},
	"GarroshBodyCheckBruteForce": {
		Name: "Brute Force",
		Text: "Enemy Heroes hit by Body Check receive 50% reduced healing for 4 seconds.",
	},
	"GarroshDecimateDeadlyCalm": {
		Name: "Deadly Calm",
		Text: "Heroes hit by Decimate deal 20% less damage for 3 seconds.",
	},
	"GarroshGroundbreakerDefensiveMeasures": {
		Name: "Defensive Measures",
		Text: "Pulling a Hero with Groundbreaker grants Garrosh a 350 Shield for 6 seconds.",
	},
	"GarroshGroundbreakerIntimidation": {
		Name: "Intimidation",
		Text: "Groundbreaker reduces the Attack Speed of Heroes hit by 40% for 4 seconds.",
	},
	"GarroshGroundbreakerMortalCombo": {
		Name: "Mortal Combo",
		Text: "If Wrecking Ball is used on a Hero within 2 seconds of pulling them with Groundbreaker, Wrecking Ball's cooldown is reduced by 10 seconds.",
	},
	"GarroshGroundbreakerRoughLanding": {
		Name: "Rough Landing",
		Text: "Heroes pulled by Groundbreaker are Slowed by 50% for 3 seconds.",
	},
	"GarroshGroundbreakerWarbreaker": {
		Name: "Warbreaker",
		Text: "Quest: Pull Heroes with Groundbreaker.\nReward: After pulling 5 Heroes, Groundbreaker deals an additional 150 damage over 3 seconds to Heroes.\nReward: After pulling 15 Heroes, reduce Groundbreaker's cooldown by 3 seconds.",
	},
	"GarroshHeroicAbilityDecimate": {
		Name: "Decimate",
		Text: "Deal 50 damage to nearby enemies and Slow them by 30% for 1.5 seconds. Deals 100% more damage to Heroes, and each Hero hit reduces the cooldown by 1 second.\nStores up to 3 charges.",
	},
	"GarroshHeroicAbilityWarlordsChallenge": {
		Name: "Warlord's Challenge",
		Text: "Silence nearby Heroes and force them to attack Garrosh for 2 seconds.",
	},
	"GarroshIndomitable": {
		Name: "Indomitable",
		Text: "Activate to become Unstoppable for 1.5 seconds.",
	},
	"GarroshOppressor": {
		Name: "Oppressor ",
		Text: "Basic Attacks against Heroes reduce the target's Spell Power by -100% for 2.5 seconds.",
	},
	"GarroshWarlordsChallengeDeathWish": {
		Name: "Death Wish",
		Text: "If an enemy Hero is killed while Taunted by Warlord's Challenge, its cooldown is reduced by 45 seconds.",
	},
	"GarroshWreckingBallEarthshaker": {
		Name: "Earthshaker",
		Text: "Wrecking Ball Stuns enemies near the impact area for 0.6 seconds.",
	},
	"GarroshWreckingBallIntotheFray": {
		Name: "Into the Fray",
		Text: "Activate to throw a nearby ally and grant them 25 Armor for 3 seconds. Deals 91 damage to nearby enemies upon impact and Slows them by 30% for 2.5 seconds.\nWhile in flight, allied Heroes are Unstoppable.",
	},
	"GarroshWreckingBallTitanicMight": {
		Name: "Titanic Might",
		Text: "Wrecking Ball now throws the 2 closest enemies.",
	},
	"GarroshWreckingBallUnrivaledStrength": {
		Name: "Unrivaled Strength",
		Text: "Increase Wrecking Ball's throw range by 25% and its damage by 125%.",
	},
	"GenericArcanePower": {
		Name: "Arcane Power",
		Text: "Activate to instantly restore 400 Mana and increase Spell Power by 15% for 10 seconds.",
	},
	"GenericDampenMagic": {
		Name: "Dampen Magic",
		Text: "Every 8 seconds, gain 50 Spell Armor against the next enemy Ability and subsequent Abilities for 1.5 seconds, reducing the damage taken by 50%.",
	},
	"GenericDoesNothing": {
		Name: "Placeholder Talent",
		Text: "This isn't the talent you're looking for.",
	},
	"GenericTalentAmplifiedHealing": {
		Name: "Amplified Healing",
		Text: "Increases regeneration effects and all healing received by 30%.",
	},
	"GenericTalentBattleMomentum": {
		Name: "Battle Momentum",
		Text: "Basic Attacks reduce Ability cooldowns by 0.5 seconds.",
	},
	"GenericTalentBerserk": {
		Name: "Berserk",
		Text: "Activate to increase your Attack Speed by 40% and Movement Speed by 10% for 4 seconds.",
	},
	"GenericTalentBlock": {
		Name: "Block",
		Text: "Every 5 seconds, gain 75 Physical Armor against the next enemy Hero Basic Attack, reducing the damage taken by 75%.\nStores up to 2 charges.",
	},
	"GenericTalentBloodForBlood": {
		Name: "Blood for Blood",
		Text: "Activate to deal 10% of target enemy Hero's Max Health and heal for twice that amount.",
	},
	"GenericTalentBribe": {
		Name: "Bribe",
		Text: "Enemy Minions killed near you grant a stack of Bribe. Use 20 stacks to bribe a Mercenary, instantly defeating them.  Does not work on Bosses.  Maximum of 100 stacks. If a camp is defeated entirely with Bribe, the camp respawns 50% faster.",
	},
	"GenericTalentBurningRage": {
		Name: "Burning Rage",
		Text: "Deal 23 damage per second to nearby enemies.",
	},
	"GenericTalentCalldownMULE": {
		Name: "Calldown: MULE",
		Text: "Activate to calldown a Mule that repairs Structures, one at a time, near target point for 40 seconds, healing for 100 Health every 1 second.  Grants 1 ammo every  3 seconds.",
	},
	"GenericTalentClairvoyance": {
		Name: "Clairvoyance",
		Text: "Activate to reveal an area for 10 seconds.  Enemies in the area are revealed for 4 seconds.",
	},
	"GenericTalentCleanse": {
		Name: "Cleanse",
		Text: "Activate to make target ally Unstoppable for 1 second. Cannot be cast on yourself.",
	},
	"GenericTalentConjurersPursuit": {
		Name: "Conjurer's Pursuit",
		Text: "Quest: Gathering Regeneration Globes increases your Mana Regeneration by 0.1 per second, up to 2.5 per second. \nReward: After gathering 25 Globes, also increase your maximum Mana by 100.",
	},
	"GenericTalentDemolitionist": {
		Name: "Demolitionist",
		Text: "Basic Attacks against Structures destroy 1 ammo and deal an additional 10% damage.",
	},
	"GenericTalentEnvenom": {
		Name: "Envenom",
		Text: "Activate to poison an enemy Hero, dealing 352 damage over 10 seconds.",
	},
	"GenericTalentFirstAid": {
		Name: "First Aid",
		Text: "Activate to heal 35.5% of your max Health over 6 seconds.",
	},
	"GenericTalentFlashoftheStorms": {
		Name: "Bolt of the Storm",
		Text: "Activate to teleport to a nearby location.",
	},
	"GenericTalentFocusedAttack": {
		Name: "Focused Attack",
		Text: "Every 10 seconds, your next Basic Attack against a Hero deals 60% additional damage. Basic Attacks reduce this cooldown by 1 second.",
	},
	"GenericTalentFollowThrough": {
		Name: "Follow Through",
		Text: "After using an Ability, your next Basic Attack within 6 seconds deals 40% additional damage.",
	},
	"GenericTalentFuryoftheStorm": {
		Name: "Fury of the Storm",
		Text: "Every 5 seconds, your next Basic Attack will deal an additional 91 damage to the target, and 228 damage to all nearby Minions, Mercenaries, and Monsters.",
	},
	"GenericTalentGatheringPower": {
		Name: "Gathering Power",
		Text: "Quest: Hero Takedowns increase Spell Power by 2%, up to 130%. This bonus Spell Power is lost on death.",
	},
	"GenericTalentGiantKiller": {
		Name: "Giant Killer",
		Text: "Basic Attacks against enemy Heroes deal bonus damage equal to 1.5% of the Hero's maximum Health.",
	},
	"GenericTalentHardenedShield": {
		Name: "Hardened Shield",
		Text: "Activate to gain 75 Armor for 4 seconds, taking 75% less damage.",
	},
	"GenericTalentHealingWard": {
		Name: "Healing Ward",
		Text: "Activate to place a ward on the ground that heals allies in an area for 1.9% of their maximum Health every second for 10 seconds.",
	},
	"GenericTalentIceBlock": {
		Name: "Ice Block",
		Text: "Activate to place yourself in Stasis and gain Invulnerability for 3 seconds.",
	},
	"GenericTalentImposingPresence": {
		Name: "Imposing Presence",
		Text: "Activate to slow the Attack Speed by 50% and Movement Speed by 20% of nearby Heroes and Summons for 2.5 seconds.\nPassive: Heroes and Summons that attack your Hero have their Attack Speed slowed by 20% for 2.5 seconds.",
	},
	"GenericTalentMercenaryLord": {
		Name: "Mercenary Lord",
		Text: "Friendly non-Boss Mercenaries near your Hero deal 50% more damage. Gain 50 Armor against Minions and Mercenaries, reducing the damage taken by 50%.",
	},
	"GenericTalentMinionKiller": {
		Name: "Minion Killer",
		Text: "Minion Killer\nDeal 25% increased damage to Minions, Mercenaries, and Summons.",
	},
	"GenericTalentNexusBlades": {
		Name: "Nexus Blades",
		Text: "Basic Attacks deal 20% more damage and Slow enemy Movement Speed by 20% for 1 second.",
	},
	"GenericTalentNexusFrenzy": {
		Name: "Nexus Frenzy",
		Text: "Increases Attack Speed by 20% and Attack Range by 1.1.",
	},
	"GenericTalentNullificationShield": {
		Name: "Nullification Shield",
		Text: "Activate to gain 60 Spell Armor for 5 seconds.",
	},
	"GenericTalentOverdrive": {
		Name: "Overdrive",
		Text: "Activate to increase Spell Power by 25% and Mana costs by 40% for 5 seconds.",
	},
	"GenericTalentPromote": {
		Name: "Promote",
		Text: "Activate to cause an allied lane Minion to take 75% reduced damage from non-Heroic targets and deal 100% bonus damage to non-Heroic targets for 30 seconds. Has 2 charges.",
	},
	"GenericTalentProtectiveShield": {
		Name: "Protective Shield",
		Text: "Activate to shield an allied Hero for 15% of their max Health for 5 seconds.",
	},
	"GenericTalentRegenerationMaster": {
		Name: "Regeneration Master",
		Text: "Quest: Gathering a Regeneration Globe increases your Health Regeneration by 1 per second, up to 30. \nReward: After gathering 30 Regeneration Globes, you also gain 500 Health.",
	},
	"GenericTalentRelentless": {
		Name: "Relentless",
		Text: "Reduces the duration of stuns, slows, and roots against your Hero by 50%.",
	},
	"GenericTalentRewind": {
		Name: "Rewind",
		Text: "Activate to reset the cooldowns of your Basic Abilities.",
	},
	"GenericTalentScoutingDrone": {
		Name: "Scouting Drone",
		Text: "Places a Scouting Drone at target location, granting vision and revealing a large area around it for 45 seconds. This drone cannot be hidden and is killed by enemies with 2 Basic Attacks.\nStores up to 2 charges.",
	},
	"GenericTalentSearingAttacks": {
		Name: "Searing Attacks",
		Text: "Activate to increase Basic Attack damage by 50% for 5 seconds. Each attack costs 15 Mana.",
	},
	"GenericTalentSeasonedMarksman": {
		Name: "Seasoned Marksman",
		Text: "Quest: Every Minion killed near you grants 0.2 Attack Damage, and Takedowns grant 0.5 Attack Damage.\nReward: Upon gaining 40 bonus Attack Damage, you can also activate Seasoned Marksman to increase your Attack Speed by 40% for 3 seconds. 60 second cooldown.",
	},
	"GenericTalentShrinkRay": {
		Name: "Shrink Ray",
		Text: "Activate to reduce an enemy Hero's damage by 50% and Movement Speed by 50% for 4 seconds.",
	},
	"GenericTalentSpellShield": {
		Name: "Spell Shield",
		Text: "Every 30 seconds, gain 50 Spell Armor against the next enemy Ability and subsequent Abilities for 3 seconds, reducing the damage taken by 50%.\nCan be toggled to allow or prevent this talent from triggering automatically.",
	},
	"GenericTalentSprint": {
		Name: "Sprint",
		Text: "Activate to gain 75% Movement Speed for 3 seconds.",
	},
	"GenericTalentStoneskin": {
		Name: "Stoneskin",
		Text: "Activate to gain 30% of your maximum Health as a Shield for 5 seconds.",
	},
	"GenericTalentStormShield": {
		Name: "Storm Shield",
		Text: "Activate to give all nearby allied Heroes a Shield for 20% of their max Health for 3 seconds.",
	},
	"GenericTalentSwiftStorm": {
		Name: "Swift Storm",
		Text: "You are no longer dismounted from taking damage.\nIncreases mount speed by 20%.",
	},
	"GenericTalentVigorousAssault": {
		Name: "Vigorous Assault",
		Text: "Basic Attacks heal for 0% of the damage dealt.",
	},
	"GenericVigorousStrikePassive": {
		Name: "Vigorous Strike",
		Text: "Basic Attacks heal for 0% of the damage dealt to the primary target.",
	},
	"GenjiCyberAgilityAgileDismount": {
		Name: "Agile Dismount",
		Text: "While mounted, the range of Cyber Agility is increased by 70% and its cooldown is reduced by 5 seconds.",
	},
	"GenjiCyberAgilityCyberShield": {
		Name: "Cyber Shield",
		Text: "Using Cyber Agility grants 50 Spell Armor for 2.5 seconds.",
	},
	"GenjiCyberAgilityDoubleJump": {
		Name: "Double Jump",
		Text: "Cyber Agility stores 2 charges but its cooldown is increased by 5 seconds.",
	},
	"GenjiCyberAgilityPathfinder": {
		Name: "Pathfinder",
		Text: "Jumping over terrain with Cyber Agility increases Genji's Movement Speed by 25% for 6 seconds. ",
	},
	"GenjiDeflectAugmentedGuard": {
		Name: "Augmented Guard",
		Text: "When Deflect ends, Genji gains a Shield equal to 50% of the damage blocked for 4 seconds.",
	},
	"GenjiDeflectDragonClaw": {
		Name: "Dragon Claw",
		Text: "After Genji blocks 330 damage with Deflect, activate to release a Dragon Claw, dealing 190 damage to all nearby enemies.",
	},
	"GenjiDeflectPerfectDefense": {
		Name: "Perfect Defense",
		Text: "Deflected damage reduces the cooldown of Deflect by 2 seconds, up to a maximum of 10 seconds.",
	},
	"GenjiDeflectReflect": {
		Name: "Reflect",
		Text: "Deflect also deals an additional 33% of the damage blocked.",
	},
	"GenjiDeflectZanshin": {
		Name: "Zanshin",
		Text: "Quest: Block 6500 damage with Deflect, including damage blocked so far.\nReward: Deflect hits all nearby enemy Heroes.",
	},
	"GenjiDodge": {
		Name: "Dodge",
		Text: "Genji dodges 1 Heroic Basic Attack every 8 seconds.",
	},
	"GenjiDragonbladeTheDragonBecomesMe": {
		Name: "The Dragon Becomes Me",
		Text: "Each time Dragonblade hits an enemy Hero, the duration of Dragonblade is increased by 0.8 seconds.",
	},
	"GenjiHeroicDragonblade": {
		Name: "Dragonblade",
		Text: "Unleash the Dragonblade for 8 seconds. While active, Dragonblade can be reactivated to lunge forward and slash in a huge arc, dealing 240 damage. If enemy Heroes are killed within 2 seconds of being hit by Dragonblade, Swift Strike's cooldown is reset.",
	},
	"GenjiHeroicXStrike": {
		Name: "X-Strike",
		Text: "Perform two slashes dealing 135 damage. The slashes detonate after 1.2 seconds causing an additional 270 damage to enemies in their area.",
	},
	"GenjiShurikenSharpenedStars": {
		Name: "Sharpened Stars",
		Text: "Shuriken now pierce all enemies hit.",
	},
	"GenjiShurikenShingan": {
		Name: "Shingan",
		Text: "Hitting an enemy with all 3 Shuriken deals 115 bonus damage.",
	},
	"GenjiShurikenShurikenMastery": {
		Name: "Shuriken Mastery",
		Text: "Quest: Hit enemy Heroes with Shuriken.\nReward: After hitting 35 Heroes, Shuriken damage is increased by 25.\nReward: After hitting 75 Heroes, Cyber Agility now refunds 2 charges of Shuriken.",
	},
	"GenjiSwiftStrikeFinalCut": {
		Name: "Final Cut",
		Text: "After 0.9 second, Swift Strike deals an additional 125 damage to all enemies in the area.",
	},
	"GenjiSwiftStrikeFlowLikeWater": {
		Name: "Flow Like Water",
		Text: "Each enemy Hero hit with Swift Strike reduces its cooldown by 3 seconds.",
	},
	"GenjiSwiftStrikeSteadyBlade": {
		Name: "Steady Blade",
		Text: "Each enemy Hero hit by Swift Strike increases the damage of the next Swift Strike by 20%. This bonus stacks up to 3 times.",
	},
	"GenjiSwiftStrikeStrikeAtTheHeart": {
		Name: "Strike At The Heart",
		Text: "Enemies hit by the end of Swift Strike take an additional 138 damage over 2 seconds.",
	},
	"GenjiSwiftStrikeSwiftAsTheWind": {
		Name: "Swift as the Wind",
		Text: "Hitting an enemy Hero with Swift Strike increases Genji's Movement Speed by 30% for 3 seconds.",
	},
	"GenjiXStrikeLivingWeapon": {
		Name: "Living Weapon",
		Text: "Each enemy Hero hit by X-Strike reduces its cooldown by 12 seconds.",
	},
	"GiantKillerTychus": {
		Name: "Problem Solver",
		Text: "Basic Attacks against enemy Heroes deal bonus damage equal to 0.5% of the Hero's maximum Health.",
	},
	"GreymaneCursedBulletGilneanRoulette": {
		Name: "Gilnean Roulette",
		Text: "Cursed Bullet now hits all enemy heroes in its path and has its cooldown reduced by 5 seconds per hit.",
	},
	"GreymaneDarkflightCyclicalNature": {
		Name: "Cyclical Nature",
		Text: "Hitting an enemy with Gilnean Cocktail reduces the cooldown of Darkflight by 3 seconds.",
	},
	"GreymaneDarkflightDisengageRunningWild": {
		Name: "Running Wild",
		Text: "Increases Darkflight and Disengage's range by 35%.",
	},
	"GreymaneDarkflightThickSkin": {
		Name: "Thick Skin",
		Text: "Using Darkflight grants 50 Physical Armor against the next 2 Hero Basic Attacks while in Worgen Form, reducing the damage taken by 50%.",
	},
	"GreymaneDisengageEyesInTheDark": {
		Name: "Eyes in the Dark",
		Text: "Disengage grants Stealth for 3 seconds.",
	},
	"GreymaneEagerWolf": {
		Name: "Eager Wolf",
		Text: "Increases the Attack Speed bonus of Inner Beast by an additional 40% after it has been active for 4 seconds.",
	},
	"GreymaneGilneanCocktailDraughtOverflow": {
		Name: "Draught Overflow",
		Text: "Increases Gilnean Cocktail's explosion length by 35%.",
	},
	"GreymaneGilneanCocktailIncendiaryElixir": {
		Name: "Incendiary Elixir",
		Text: "Quest: Every time Greymane hits an enemy Hero with the explosion damage from Gilnean Cocktail, increase its explosion damage by 15, up to 225.\nReward: After hitting 15 Heroes, the cooldown is also reduced by 4 seconds.",
	},
	"GreymaneGilneanCocktailPerfectAim": {
		Name: "Perfect Aim",
		Text: "Increases Gilnean Cocktail's range by 30% and refunds 50 Mana if it hits an enemy Hero.",
	},
	"GreymaneGoForTheThroatUnleashed": {
		Name: "Unleashed",
		Text: "If the free cast of Go for the Throat kills its target, another free cast is provided that deals 35% more damage, up to 140%.",
	},
	"GreymaneHeroicAbilityCursedBullet": {
		Name: "Cursed Bullet",
		Text: "Greymane shapeshifts into a Human and fires a bullet that hits the first enemy Hero in its path, dealing 35% of their current Health in damage.\nDoes not affect Vehicles.",
	},
	"GreymaneHeroicAbilityGoForTheThroat": {
		Name: "Go for the Throat",
		Text: "Leap at an enemy Hero and shapeshift into a Worgen, slashing for 355 damage. If this kills them, the Ability can be used a second time within 10 seconds for free.",
	},
	"GreymaneHuntersBlunderbuss": {
		Name: "Hunter's Blunderbuss",
		Text: "Human Basic Attacks splash for 100% damage behind the target.",
	},
	"GreymaneInnerBeastInsatiable": {
		Name: "Insatiable",
		Text: "Inner Beast causes Basic Attacks to restore 10 Mana.",
	},
	"GreymaneInnerBeastOnTheProwl": {
		Name: "On the Prowl",
		Text: "Inner Beast increases Movement Speed by 30% once it has been active for 3 seconds.",
	},
	"GreymaneInnerBeastViciousness": {
		Name: "Viciousness",
		Text: "Increases Inner Beast's duration to 4 seconds, and causes Ability damage to also refresh its duration.",
	},
	"GreymaneInnerBeastWolfheart": {
		Name: "Wolfheart",
		Text: "Increase the cooldown reduction from Basic Attacks during Inner Beast from 0.5 to 1 second.",
	},
	"GreymaneRazorSwipeUnfetteredAssault": {
		Name: "Unfettered Assault",
		Text: "Increases Razor Swipe's lunge distance by 60%.",
	},
	"GreymaneRazorSwipeVisceralAttacks": {
		Name: "Visceral Attacks",
		Text: "Worgen Basic Attacks reduce Razor Swipe's cooldown by 1 second.",
	},
	"GreymaneToothAndClaw": {
		Name: "Tooth and Claw",
		Text: "Worgen Basic Attacks cleave for 100% damage.",
	},
	"GreymaneWizenedDuelist": {
		Name: "Wizened Duelist",
		Text: "Quest: Hero takedowns increase Basic Attack damage by 3%, up to 30%. This bonus is lost on death.",
	},
	"GreymaneWorgenFormAlphaKiller": {
		Name: "Alpha Killer",
		Text: "Worgen Basic Attacks against Heroes deal bonus damage equal to 3% of the Hero's maximum Health.",
	},
	"GreymaneWorgenFormQuicksilverBullets": {
		Name: "Quicksilver Bullets",
		Text: "Increases Greymane's Human Basic Attack range by 1.1.",
	},
	"GuldanChaoticEnergy": {
		Name: "Chaotic Energy",
		Text: "Quest: Gathering a Regen Globe causes Gul'dan's next Basic Ability with a Mana cost to refund 40 Mana\nReward: After Gathering 20 Regen Globes Gul'dan's Basic Ability Mana costs are permanently reduced by 20.",
	},
	"GuldanConsumeSoul": {
		Name: "Consume Soul",
		Text: "Instantly kill an enemy Minion and heal for 365 Health.\nStores up to 2 charges.",
	},
	"GuldanCorruptionCurseOfExhaustion": {
		Name: "Curse of Exhaustion",
		Text: "Upon expiration, Corruption slows enemy Movement Speed by 50% for 2.5 seconds.",
	},
	"GuldanCorruptionDemonicSight": {
		Name: "Demonic Sight",
		Text: "Corruption reveals enemies for the duration.",
	},
	"GuldanCorruptionEchoedCorruption": {
		Name: "Echoed Corruption",
		Text: "Quest: Hit 40 enemy Heroes with Corruption.\nReward: After the third strike, Corruption strikes 3 times in the opposite direction.",
	},
	"GuldanCorruptionRuinousAffliction": {
		Name: "Ruinous Affliction",
		Text: "Corruption deals an additional 78 damage on impact. If the enemy is hit by all three strikes, the third strike deals 233 damage.",
	},
	"GuldanDarkBargain": {
		Name: "Dark Bargain",
		Text: "Gul'dan's maximum Health is permanently increased by 25% but his respawn time is increased by 15 seconds.",
	},
	"GuldanDemonicCircle": {
		Name: "Demonic Circle",
		Text: "Summon a Demonic Circle at Gul'dan's location. Activate to teleport Gul'dan to the Demonic Circle.",
	},
	"GuldanDrainLifeDevourTheFrail": {
		Name: "Devour the Frail",
		Text: "Drain Life deals 50% more damage to enemies below 50% Health.",
	},
	"GuldanDrainLifeGlyphOfDrainLife": {
		Name: "Glyph of Drain Life",
		Text: "Increases the cast range of Drain Life by 25%.",
	},
	"GuldanDrainLifeHarvestLife": {
		Name: "Harvest Life",
		Text: "Drain Life heals for 50% more Health when used on Heroes.",
	},
	"GuldanDrainLifeHealthFunnel": {
		Name: "Health Funnel",
		Text: "If an enemy dies while under the effect of Drain Life, the cooldown is instantly refreshed.",
	},
	"GuldanFelFlameBoundByShadow": {
		Name: "Bound by Shadow",
		Text: "Each enemy Hero hit with Fel Flame reduces the cooldown of Corruption by 1 second.",
	},
	"GuldanFelFlameFelArmor": {
		Name: "Fel Armor",
		Text: "Hitting an enemy Hero with Fel Flame grants 40 Spell Armor for 2.5 seconds, reducing Ability Damage taken by 40%. ",
	},
	"GuldanFelFlamePursuitOfFlame": {
		Name: "Pursuit of Flame",
		Text: "Quest: Hit 40 enemy Heroes with Fel Flame.\nReward: Fel Flame's radius is increased by 10%.",
	},
	"GuldanFelFlameRampantHellfire": {
		Name: "Rampant Hellfire",
		Text: "Fel Flame's damage is increased by 8% for 5 seconds when hitting an enemy Hero. This can stack up to 5 times.",
	},
	"GuldanHealthstone": {
		Name: "Healthstone",
		Text: "Activate to heal for 25% of Gul'dan's maximum Health.",
	},
	"GuldanHorrify": {
		Name: "Horrify",
		Text: "After 0.5 seconds, deal 120 damage to enemy Heroes in an area and Fear them for 2 seconds. While Feared, Heroes are Silenced and are forced to run away from Horrify's center.",
	},
	"GuldanHorrifyHaunt": {
		Name: "Haunt",
		Text: "Increases the duration of Horrify by 1 second, and while Feared, enemies lose 25 Armor, causing them to take 25% increased damage.",
	},
	"GuldanLifeTapDarknessWithin": {
		Name: "Darkness Within",
		Text: "After using Life Tap, the next Ability Gul'dan casts has 25% increased Spell Power.",
	},
	"GuldanLifeTapHungerforPower": {
		Name: "Hunger for Power",
		Text: "Increases Spell Power by 15% but reduces healing received from allies by 25%.",
	},
	"GuldanLifeTapImprovedLifeTap": {
		Name: "Improved Life Tap",
		Text: "Life Tap now restores 35% Mana",
	},
	"GuldanRainOfDestruction": {
		Name: "Rain of Destruction",
		Text: "Channel to summon a rain of meteors in an area over 7 seconds. Each meteor deals 165 damage in a small area.",
	},
	"GuldanRainOfDestructionDeepImpact": {
		Name: "Deep Impact",
		Text: "Rain of Destruction slows enemies Movement Speed by 90% for 0.5 seconds.",
	},
	"HeroGenericExecutionerPassive": {
		Name: "Executioner",
		Text: "Attacking a Hero that is slowed, rooted, or stunned increases your Basic Attack damage by 30% for 3 seconds.",
	},
	"IllidanBladesOfAzzinoth": {
		Name: "Blades of Azzinoth",
		Text: "Hitting 5 Heroes with Sweeping Strike allows Blades of Azzinoth to be activated, increasing Basic Attack damage by 75% for 8 seconds.",
	},
	"IllidanCombatStyleHuntersOnslaught": {
		Name: "Hunter's Onslaught",
		Text: "Basic Abilities heal for 25% of the damage they deal. This bonus is doubled versus Heroes.",
	},
	"IllidanCombatStyleThrillOfBattle": {
		Name: "Thrill of Battle",
		Text: "Activate to double the cooldown reduction from Basic Attacks for 8 seconds.",
	},
	"IllidanElusiveStrike": {
		Name: "Elusive Strike",
		Text: "Sweeping Strike reduces the cooldown of Evasion by 3 seconds every time it damages an enemy Hero.",
	},
	"IllidanFieryBrand": {
		Name: "Fiery Brand",
		Text: "Every 4th attack against the same Hero deals an additional 6% of their maximum Health as damage.",
	},
	"IllidanHeroicAbilityMetamorphosis": {
		Name: "Metamorphosis",
		Text: "Transform into Demon Form at the target location, dealing 46 damage in the area. Temporarily increases maximum Health by 200 for each Hero hit by the initial impact. Lasts for 18 seconds.",
	},
	"IllidanHeroicAbilityTheHunt": {
		Name: "The Hunt",
		Text: "Charge to target unit, dealing 251 damage on impact and stunning for 1 second.",
	},
	"IllidanMasteryBatteredAssaultSweepingStrike": {
		Name: "Battered Assault",
		Text: "Increase Sweeping Strike's Basic Attack damage bonus duration from 3 to 5 seconds. If Sweeping Strike hits 2 Heroes its damage bonus is increased from 35% to 125%.",
	},
	"IllidanMasteryDemonicFormMetamorphosis": {
		Name: "Demonic Form",
		Text: "Permanently remain in Demonic Form. Metamorphosis also increases Attack Speed by 20% and reduces the duration of slows, roots, and stuns by 50%. Illidan can now mount in Demonic Form.",
	},
	"IllidanMasteryFriendOrFoeDive": {
		Name: "Friend or Foe",
		Text: "Increases the range of Dive by 20% and allows it to be used to dive to allied Heroes.",
	},
	"IllidanMasteryImmolationSweepingStrike": {
		Name: "Immolation",
		Text: "After using Sweeping Strike, burn nearby enemies for 22 damage a second for 4 seconds.",
	},
	"IllidanMasteryLungeDive": {
		Name: "Lunge",
		Text: "Increases the range of Dive by 30%.",
	},
	"IllidanMasteryMarkedforDeathDive": {
		Name: "Marked for Death",
		Text: "Dive deals an extra 180 damage if used consecutively on a target within 10 seconds.",
	},
	"IllidanMasteryNowhereToHideTheHunt": {
		Name: "Nowhere to Hide",
		Text: "The Hunt gains unlimited range. Illidan passively reveals enemy Heroes below 25% Health anywhere on the Battleground.",
	},
	"IllidanMasteryRapidChaseDive": {
		Name: "Rapid Chase",
		Text: "Dive grants 15% Movement Speed for 2.2 seconds.",
	},
	"IllidanMasterySecondSweepSweepingStrike": {
		Name: "Second Sweep",
		Text: "Store up to 2 charges of Sweeping Strike.",
	},
	"IllidanMasteryShadowShieldEvasion": {
		Name: "Shadow Shield",
		Text: "Evasion grants a 0 point Shield for 0 seconds.",
	},
	"IllidanMasterySixthSenseEvasion": {
		Name: "Sixth Sense",
		Text: "While active, Evasion grants Illidan 75 Spell Armor against the next 2 sources of Spell Damage, reducing their damage by 75%.",
	},
	"IllidanMasteryUnboundSweepingStrike": {
		Name: "Unbound",
		Text: "Sweeping Strike can go over walls and terrain.\nReward: After hitting 30 Heroes with Sweeping Strike, gain a second charge.",
	},
	"IllidanNimbleDefender": {
		Name: "Nimble Defender",
		Text: "If Sweeping Strikes hits an enemy Hero, gain 25 Armor for 2 seconds, reducing damage taken by 25%.",
	},
	"IllidanReflexiveBlock": {
		Name: "Reflexive Block",
		Text: "Using Dive grants 75 Physical Armor against the next 2 Hero Basic Attacks for 3 seconds, reducing the damage taken by 75%.",
	},
	"IllidanSweepingStrikesThirstingBlade": {
		Name: "Thirsting Blade",
		Text: "Betrayer's Thirst's healing from Basic Attacks is increased from 30% to 50% while Sweeping Strike's damage bonus is active.",
	},
	"IllidanUnendingHatredPassive": {
		Name: "Unending Hatred",
		Text: "Quest: Minion kills grant 0.2 Basic Attack Damage. Hero Takedowns grant 1 Basic Attack Damage.\nReward: After gaining 20 increased Basic Attack Damage, receive an additional 20 Basic Attack Damage.",
	},
	"JainaBlizzardSnowstorm": {
		Name: "Snowstorm",
		Text: "Increase the radius of Blizzard by 30%.",
	},
	"JainaBlizzardStormFront": {
		Name: "Storm Front",
		Text: "Increase cast range of Blizzard by 100%.",
	},
	"JainaConeOfColdIceFloes": {
		Name: "Ice Floes",
		Text: "Increase the width of Cone of Cold by 100% and cause each chilled Hero hit to reduce its cooldown by 2.5 seconds, up to a maximum of 5 seconds per use.",
	},
	"JainaConeOfColdNorthernExposure": {
		Name: "Northern Exposure",
		Text: "Enemies damaged by Cone of Cold have their Armor lowered by 25 for 2 seconds, increasing their damage taken by 25%.",
	},
	"JainaConeOfColdNumbingBlast": {
		Name: "Numbing Blast",
		Text: "Cone of Cold Roots Chilled targets for 1 second.",
	},
	"JainaFingersOfFrost": {
		Name: "Fingers of Frost",
		Text: "Quest: Gathering Regeneration Globes to increase Jaina's Mana regeneration by 0.1 per second, up to 2 per second. \nReward: After gathering 20 Globes, the damage bonus from Frostbite is increased by 10%.",
	},
	"JainaFrigidTransmission": {
		Name: "Ice Blink",
		Text: "Activate to teleport a short distance and Chill all nearby enemies.",
	},
	"JainaFrostbiteArcaneIntellect": {
		Name: "Arcane Intellect",
		Text: "Dealing damage to a Chilled target returns Mana to Jaina. Basic Attacks return 0 Mana and Abilities return 0.",
	},
	"JainaFrostbiteDeepChill": {
		Name: "Deep Chill",
		Text: "Frostbite's Chill can now stack a second time, slowing targets up to 40%.",
	},
	"JainaFrostbiteFrostArmor": {
		Name: "Frost Armor",
		Text: "Enemies that attack Jaina are Chilled. Additionally, every 10 seconds, Jaina gains 50 Physical Armor against the next enemy Hero Basic Attack, reducing the damage taken by 50%.",
	},
	"JainaFrostbiteFrostbitten": {
		Name: "Frostbitten",
		Text: "Increases the damage bonus of Frostbite from 50% to 65%.",
	},
	"JainaFrostbiteIceBarrier": {
		Name: "Ice Barrier",
		Text: "Gain a Shield for 25% of Ability damage dealt to Chilled targets. The shield lasts 4 seconds.",
	},
	"JainaFrostbiteLingeringChill": {
		Name: "Lingering Chill",
		Text: "Increase the duration of Chill from 4 seconds to 6 seconds.",
	},
	"JainaFrostboltFrostShards": {
		Name: "Frost Shards",
		Text: "Frostbolt will now pierce the first target to hit an additional target behind them.",
	},
	"JainaFrostboltIceLance": {
		Name: "Ice Lance",
		Text: "Hitting a Chilled target with Frostbolt reduces its cooldown by 2 seconds and restores 10 Mana.",
	},
	"JainaFrostboltWintersReach": {
		Name: "Winter's Reach",
		Text: "Increase the range of Frostbolt by 30%.",
	},
	"JainaHeroicRingOfFrost": {
		Name: "Ring of Frost",
		Text: "After a 1.5 second delay, create a Ring of Frost in an area that deals 310 damage and Roots enemies for 3 seconds. The ring persists for 3 seconds afterward, Chilling any enemies who touch it.",
	},
	"JainaHeroicSummonWaterElemental": {
		Name: "Summon Water Elemental",
		Text: "Summon a Water Elemental at target location. The Water Elemental's Basic Attacks deal 78 damage, splash for 25% damage and Chill. The Ability can be reactivated to retarget the Water Elemental.  Lasts 20 seconds.",
	},
	"JainaIcefuryWand": {
		Name: "Icefury Wand",
		Text: "Basic Attacks against Chilled Heroes deal 75% more damage and lower the cooldown of Blizzard by 1 second.",
	},
	"JainaIcyVeins": {
		Name: "Icy Veins",
		Text: "Activate to make Jaina's Basic Abilities' cooldowns recharge 200% faster and reduce their Mana cost by 50% for 5 seconds.",
	},
	"JainaImprovedIceBlock": {
		Name: "Improved Ice Block",
		Text: "Activate to place Jaina in Stasis and gain Invulnerability for 2.5 seconds. When this effect expires, nearby enemies are Chilled.",
	},
	"JainaRingOfFrostColdSnap": {
		Name: "Cold Snap",
		Text: "The center area of Ring of Frost deals 310 damage and Roots enemies after the outer ring expires. Each enemy Hero hit reduces its cooldown by 10 seconds.",
	},
	"JainaSummonWaterElementalWintermute": {
		Name: "Wintermute",
		Text: "Increases the cast range of Water Elemental by 50%, and the Water Elemental will now mimic Jaina's Basic Abilities for 50% damage.",
	},
	"JunkratBombPacks": {
		Name: "Bomb Packs",
		Text: "Activate to place a Bomb Pack on the ground that can be picked up by an allied Hero. Heroes that die while in possession of a Bomb Pack drop 3 grenades that explode after 1.2 seconds, dealing 250 damage in an area.\nLimit 0 Bomb Pack per Hero, and 5 Bomb Packs on the ground at one time.",
	},
	"JunkratConcussionMineBOOMPOW": {
		Name: "BOOM POW",
		Text: "Hitting an enemy Hero with Concussion Mine reduces its cooldown by 10 seconds.",
	},
	"JunkratConcussionMineBoggedDown": {
		Name: "Bogged Down",
		Text: "Enemies launched by Concussion Mine are Slowed by 60% for 2 seconds upon landing.",
	},
	"JunkratConcussionMineBonzerHits": {
		Name: "Bonzer Hits",
		Text: "Quest: Hit enemy Heroes 12 times with Concussion Mine.\nReward: Concussion Mine deals 40% more damage and knocks enemies back 30% farther.",
	},
	"JunkratConcussionMineRipperAir": {
		Name: "Ripper Air",
		Text: "Junkrat is knocked back 50% farther by Concussion Mine. Additionally, Concussion Mine's cooldown is reduced by 12 seconds if only Junkrat is hit.",
	},
	"JunkratFragLauncherBouncyBouncy": {
		Name: "Bouncy Bouncy",
		Text: "Upon colliding with terrain for the first time, Frag Launcher grenades deal 20% more damage.",
	},
	"JunkratFragLauncherBurstFire": {
		Name: "Burst Fire",
		Text: "Frag Launcher fires all of its charges in a continuous burst, its cooldown is reduced by 6 seconds, but it loses 1 maximum charge.",
	},
	"JunkratFragLauncherCannonball": {
		Name: "Cannonball!",
		Text: "Increase the radius and explosion radius of grenades from Basic Attacks and Frag Launcher by 50%.",
	},
	"JunkratFragLauncherEndlessNades": {
		Name: "Endless Nades",
		Text: "Hitting an enemy Hero with Frag Launcher reduces its cooldown by 1.5 seconds.",
	},
	"JunkratFragLauncherExtraWoundTimers": {
		Name: "Extra-Wound Timers",
		Text: "Frag Launcher grenades last an additional 2 seconds before automatically detonating.",
	},
	"JunkratFragLauncherPutSomeEnglishOnIt": {
		Name: "Put Some English On It",
		Text: "Increase Frag Launcher's travel distance by 50.4%, but does not increase its speed.",
	},
	"JunkratFragLauncherTasteForExplosions": {
		Name: "Taste For Explosions",
		Text: "Quest: Hitting a Hero with Frag Launcher increases its damage by 0.5, up to 100.",
	},
	"JunkratFragLauncherTrickyShuffles": {
		Name: "Tricky Shuffles",
		Text: "While Frag Launcher has no charges left, gain 15% Movement Speed.",
	},
	"JunkratHeroicAbilityRIPTire": {
		Name: "RIP-Tire",
		Text: "Create a motorized bomb with 530 Health that lasts 15 seconds. While active, Junkrat is immobile but gains control of RIP-Tire's movement.\nRIP-Tire can be reactivated to detonate immediately, knocking nearby enemies back and dealing 475, 625, or 775 damage to enemies depending on how close they are to the center of the blast (with the highest amount near the center).",
	},
	"JunkratHeroicRocketRide": {
		Name: "Rocket Ride",
		Text: "After 1.2 seconds, Junkrat launches into the air. While in the air, he can steer the landing location by moving.\nAfter 3.5 seconds, Junkrat lands, dealing 750 damage to nearby enemies and activating Total Mayhem. 5 seconds after landing, Junkrat reappears at the Altar and gains 150% additional Movement Speed until dismounted.",
	},
	"JunkratRIPTireExtraOomph": {
		Name: "Extra Oomph",
		Text: "Increase RIP-Tire's knockback distance by 50%, and its cooldown is reduced by 20 seconds for each enemy Hero hit.",
	},
	"JunkratRocketRidePuckishScamp": {
		Name: "Puckish Scamp",
		Text: "Reduce Rocket Ride's cooldown by 70 seconds.",
	},
	"JunkratSpreadVolley": {
		Name: "Spread Volley",
		Text: "Activate to make Frag Launcher fire 2 additional grenades in a spread. Works for up to 4 total charges, or until Frag Launcher runs out of charges.",
	},
	"JunkratSteelTrapBigAs": {
		Name: "Big As",
		Text: "Increase Steel Trap's radius and damage by 50%.",
	},
	"JunkratSteelTrapChatteringTeeth": {
		Name: "Chattering Teeth",
		Text: "Steel Traps will slowly chase nearby enemy Heroes.",
	},
	"JunkratSteelTrapGottaTrapEmAll": {
		Name: "Gotta Trap 'Em All!",
		Text: "Quest: Hit 7 Heroes with Steel Trap.\nReward: Reduce Steel Trap's cooldown by 3 seconds and increase the maximum number of active traps to 3.",
	},
	"JunkratSteelTrapStickyWicket": {
		Name: "Sticky Wicket",
		Text: "Steel Trap no longer Roots enemies, and instead Slows them by 90% for 3.5 seconds.",
	},
	"KaelthasArcaneDynamo": {
		Name: "Arcane Dynamo",
		Text: "Using a Basic Ability increases Kael'thas's Spell Power by 3% for 5 seconds, stacking up to 15%.",
	},
	"KaelthasFlamestrikeBurnedFlesh": {
		Name: "Burned Flesh",
		Text: "When Flamestrike damages 2 or more Heroes, they take additional damage equal to 8% of their maximum Health.",
	},
	"KaelthasFlamestrikeConvection": {
		Name: "Convection",
		Text: "Quest: Damaging an enemy Hero with Flamestrike increase its damage by 5, up to 100. This bonus is lost on death.\nReward: After hitting 20 enemy Heroes, increase Flamestrike damage by an additional 100 and no longer lose any bonuses on death.",
	},
	"KaelthasFlamestrikeFuryOfTheSunwell": {
		Name: "Fury of the Sunwell",
		Text: "Flamestrike will explode a second time 1.5 seconds later.",
	},
	"KaelthasFlamestrikeManaTap": {
		Name: "Mana Tap",
		Text: "Activating Verdant Spheres restores 3.9% of Kael'thas's maximum Mana.",
	},
	"KaelthasGravityLapseEnergyRoil": {
		Name: "Energy Roil",
		Text: "When Gravity Lapse hits a Hero, reduce its cooldown by 9 seconds.",
	},
	"KaelthasGravityLapseGravityThrow": {
		Name: "Gravity Throw",
		Text: "Increases the duration of Gravity Lapse's stun by 33% and causes it to instantly destroy Minions.",
	},
	"KaelthasGravityLapseNetherWind": {
		Name: "Nether Wind",
		Text: "Increases Gravity Lapse's range by 30%. When Gravity Lapse hits an enemy Hero, refund 80 Mana.",
	},
	"KaelthasGravityLapseTriOptimal": {
		Name: "Tri-Optimal",
		Text: "The cooldown of Verdant Spheres is refreshed by 2 seconds per target hit with Gravity Lapse.",
	},
	"KaelthasHeroicAbilityPhoenix": {
		Name: "Phoenix",
		Text: "Launch a Phoenix to an area, dealing 78 damage to enemies along the way. The Phoenix persists for 7 seconds, attacking enemies for 78 damage and splashing for 50%.",
	},
	"KaelthasHeroicAbilityPyroblast": {
		Name: "Pyroblast",
		Text: "After 1.5 seconds, cast a slow-moving fireball that deals 810 damage to an enemy Hero and 405 damage to enemies nearby.",
	},
	"KaelthasLivingBombBackdraft": {
		Name: "Backdraft",
		Text: "Living Bomb's explosion slows enemy Movement Speed by 30% for 2 seconds.",
	},
	"KaelthasLivingBombChainBomb": {
		Name: "Chain Bomb",
		Text: "Living Bomb's explosion applies Living Bomb to the 3 closest enemies not already affected by it.  Prefers Heroes.",
	},
	"KaelthasLivingBombFissionBomb": {
		Name: "Fission Bomb",
		Text: "Increases Living Bomb's explosion radius by 20%.",
	},
	"KaelthasLivingBombPyromaniac": {
		Name: "Pyromaniac",
		Text: "Each time Living Bomb deals periodic damage, Kael'thas's Basic Ability cooldowns are reduced by 1 second.",
	},
	"KaelthasLivingBombSunKingsFury": {
		Name: "Sun King's Fury",
		Text: "Living Bombs that spread to enemy Heroes deal 35% more damage.",
	},
	"KaelthasManaAddict": {
		Name: "Mana Addict",
		Text: "Quest: Gathering a Regeneration Globe increases Kael'thas's maximum Mana by 15.\nReward: After gathering 20 Globes, Kael'thas can activate Arcane Barrier to gain a Shield equal to 100% of his maximum Mana for 4 seconds. 45 second cooldown.",
	},
	"KaelthasMasterOfFlames": {
		Name: "Master of Flames",
		Text: "Living Bomb's spread from explosions can now also spread Living Bomb.",
	},
	"KaelthasMasteryFlamethrower": {
		Name: "Flamethrower",
		Text: "Increases the cast range of Flamestrike by 40%. If Flamestrike hits 2 or more Heroes, its cooldown is reduced by 4 seconds.",
	},
	"KaelthasMasterySunfireEnchantment": {
		Name: "Sunfire Enchantment",
		Text: "Activating Verdant Spheres causes Kael'thas's next Basic Attack to instead deal 165 Spell Damage.",
	},
	"KaelthasPhoenixRebirth": {
		Name: "Rebirth",
		Text: "Increases Phoenix duration by 100%. While alive, the Phoenix may be ordered to move to a different location.",
	},
	"KaelthasPyroblastPresenceOfMind": {
		Name: "Presence Of Mind",
		Text: "Increases Pyroblast's explosion radius by 50% and reduces its cooldown by 25 seconds per enemy Hero hit.",
	},
	"KaelthasTwinSpheres": {
		Name: "Twin Spheres",
		Text: "Verdant Spheres gains a second charge.",
	},
	"KaelthasVerdantSpheresFelInfusion": {
		Name: "Fel Infusion",
		Text: "Increases Kael'thas's Spell Power by 4%. Kael'thas heals for 94 Health when activating Verdant Spheres.",
	},
	"KelThuzadAcceleratedDecay": {
		Name: "Accelerated Decay",
		Text: "Each time a Hero is hit by Death and Decay's pool, they take 25% more periodic damage from Death and Decay for the next 4 seconds, stacking up to 6 times.",
	},
	"KelThuzadArcaneEchoes": {
		Name: "Arcane Echoes",
		Text: "Whenever Kel'Thuzad hits an enemy Hero with Death and Decay's explosion, its cooldown is reduced by 1.2 seconds.",
	},
	"KelThuzadArchlichArmor": {
		Name: "Armor of the Archlich",
		Text: "Activate to gain 50 Physical Armor for 4 seconds. Upon activation, nearby enemies take 45 damage and are Slowed by 35% for 4 seconds.",
	},
	"KelThuzadBarbedChains": {
		Name: "Barbed Chains",
		Text: "Increase Chains of Kel'Thuzad damage by 125%. After gaining 30 Blight, Chains of Kel'Thuzad reduces the Armor of pulled Heroes by 15 for 4 seconds.",
	},
	"KelThuzadBlightedFrost": {
		Name: "Blighted Frost",
		Text: "Frost Nova deals 50% more damage to enemies in the center. After gaining 30 Blight, increase Frost Nova's Root duration by 0.5 seconds.",
	},
	"KelThuzadChainLink": {
		Name: "Chain-Link",
		Text: "Pulling 2 Heroes together with Chains of Kel'Thuzad reduces its cooldown by 4 seconds and refunds its Mana cost.",
	},
	"KelThuzadChainsOfIce": {
		Name: "Chains of Ice",
		Text: "After Chains of Kel'Thuzad's Stun expires, the enemy is also Slowed by 60% for 1 second.",
	},
	"KelThuzadChillingTouch": {
		Name: "Chilling Touch",
		Text: "Every 8 seconds, Kel'Thuzad's next Basic Attack hits all enemies in the area for 100%  additional damage, deals Spell damage instead of Physical, and Slows by 30% for 2 seconds.",
	},
	"KelThuzadDeathchill": {
		Name: "Deathchill",
		Text: "When Frost Blast's Root expires, enemies are Slowed by 40% for 3 seconds. Heroes killed while under the effects of Frost Blast instantly release another Frost Blast explosion.",
	},
	"KelThuzadFrozenTomb": {
		Name: "Frost Blast",
		Text: "Launch a meteor of ice at an enemy Hero. Upon impact, the meteor deals 100 damage to its target and 275 damage to enemies in the area. All enemies hit by Frost Blast are Rooted for 1.5 seconds.",
	},
	"KelThuzadGlacialSpike": {
		Name: "Glacial Spike",
		Text: "Activate to create a spike that detonates after 4 seconds, dealing 60 damage to nearby enemies. The spike can be affected by Chains of Kel'Thuzad.",
	},
	"KelThuzadHungeringCold": {
		Name: "Hungering Cold",
		Text: "Enemies Rooted by Frost Nova take an additional 55 damage each time they are damaged by Kel'Thuzad during the next 4 seconds.",
	},
	"KelThuzadIcyGrasp": {
		Name: "Icy Grasp",
		Text: "Increase Frost Nova's Slow duration by 2 seconds.",
	},
	"KelThuzadMasterOfTheColdDark": {
		Name: "Master of the Cold Dark",
		Text: "Kel'Thuzad becomes increasingly stronger as he disrupts enemies",
	},
	"KelThuzadMightOfTheScourge": {
		Name: "Might of the Scourge",
		Text: "Hitting a Hero with Shadow Fissure resets its cooldown.",
	},
	"KelThuzadPhylactery": {
		Name: "Phylactery of Kel'Thuzad",
		Text: "Quest: Collect 12 Regeneration Globes to charge Kel'Thuzad's Phylactery.\nReward: Kel'Thuzad's Phylactery can be activated while dead to immediately respawn at the Hall of Storms, but must be charged again.\nPassive: Kel'Thuzad heals for 10% of all Spell Damage dealt while the Phylactery is charged.",
	},
	"KelThuzadPowerOfIcecrown": {
		Name: "Power of Icecrown",
		Text: "Stunning, Rooting, or Slowing a Hero grants 6% Spell Power for 10 seconds, stacking up to 5 times.",
	},
	"KelThuzadShadowFissure": {
		Name: "Shadow Fissure",
		Text: "Create a fissure anywhere on the Battleground that explodes after 1.5 seconds, dealing 400 damage to enemy Heroes in its area.",
	},
	"KelThuzadShiftingMalice": {
		Name: "Shifting Malice",
		Text: "Activate to dash forward, dealing 150 damage to enemies in the path. Takedowns reset the cooldown of Shifting Malice.",
	},
	"KelThuzadStripShields": {
		Name: "Strip Shields",
		Text: "Pulling a Hero with Chains of Kel'Thuzad grants Kel'Thuzad a permanent 120 Shield, stacking up to 2 times.\nAdditionally, Chains of Kel'Thuzad deals up to 270 bonus damage to Shields.",
	},
	"KelThuzadTheDamnedReturn": {
		Name: "The Damned Return",
		Text: "Activate to create a Shade of Naxxramas that lasts 15 seconds and mimics Kel'Thuzad's casts of Death and Decay.",
	},
	"KelThuzadThePlaguelands": {
		Name: "The Plaguelands",
		Text: "Increase Death and Decay's duration by 1 second. After gaining 30 Blight, increase the radius of Death and Decay's pool by 30%.",
	},
	"KerriganAssimilationMastery": {
		Name: "Assimilation Mastery",
		Text: "Increases the duration of Assimilation by 100%. While Assimilation is active, Kerrigan's Health and Mana regeneration is increased by 100%.",
	},
	"KerriganCombatStyleDoubleStrike": {
		Name: "Double Strike",
		Text: "When a Basic Ability damages an enemy, Kerrigan's next Basic Attack hits for 75% bonus damage.",
	},
	"KerriganCombatStyleFuryoftheSwarm": {
		Name: "Fury of the Swarm",
		Text: "Kerrigan's Basic Attacks splash for 60% damage around the target.",
	},
	"KerriganCombatStyleLingeringEssence": {
		Name: "Lingering Essence",
		Text: "Increases Assimilation Shield's duration to 20 seconds.",
	},
	"KerriganEssenceForEssence": {
		Name: "Essence for Essence",
		Text: "Activate to deal 10% of target enemy Hero's Max Health and gain Assimilation Shields for twice that amount.",
	},
	"KerriganHeroicAbilityMaelstrom": {
		Name: "Maelstrom",
		Text: "Deals 74 damage per second to nearby enemies. Lasts for 7 seconds.",
	},
	"KerriganHeroicAbilitySummonUltralisk": {
		Name: "Summon Ultralisk",
		Text: "Summon an Ultralisk that attacks the target to deal 100 damage. Attacks splash to nearby enemies for 50% damage. Can reactivate the Ability to retarget the Ultralisk. Lasts for 20 seconds.",
	},
	"KerriganMasteryAdaptation": {
		Name: "Adaptation",
		Text: "Ravage can be used to jump to allies, refunding 50% of the cooldown and 100% of the Mana cost.",
	},
	"KerriganMasteryAggressiveDefense": {
		Name: "Aggressive Defense",
		Text: "Increases base Shield amount gained from Assimilation by 100%.",
	},
	"KerriganMasteryCrushingSwarm": {
		Name: "Impaling Swarm",
		Text: "Impaling Blades spawns 2 Zerglings.",
	},
	"KerriganMasteryEviscerate": {
		Name: "Eviscerate",
		Text: "Increases Ravage's range by 40%.",
	},
	"KerriganMasteryImpalingBladesBladeTorrent": {
		Name: "Blade Torrent",
		Text: "Increases Impaling Blades' radius by 30%.",
	},
	"KerriganMasteryImpalingBladesSharpenedBlades": {
		Name: "Sharpened Blades",
		Text: "Impaling Blades deals 25% more damage.",
	},
	"KerriganMasteryMaelstromOmegastorm": {
		Name: "Omegastorm",
		Text: "Maelstrom size increased by 25% . Amount of Assimilation Shields generated by Maelstrom increased by 100%.",
	},
	"KerriganMasteryPrimalGraspEnergizingGrasp": {
		Name: "Energizing Grasp",
		Text: "Reduces the Mana cost of Primal Grasp by 75%. ",
	},
	"KerriganMasteryPrimalGraspPsionicPulse": {
		Name: "Psionic Pulse",
		Text: "After casting Primal Grasp, deal 36 damage per second to nearby enemies. Lasts 5 seconds.",
	},
	"KerriganMasteryRavageCleanKill": {
		Name: "Clean Kill",
		Text: "If Ravage kills the target, it restores 75% of its Mana cost and increases the damage of the next Ravage by 25%.",
	},
	"KerriganMasteryRavageSiphoningImpact": {
		Name: "Siphoning Impact",
		Text: "If the targeted enemy dies within 1.5 seconds, Kerrigan heals for 10.2% of her maximum Health.",
	},
	"KerriganMasteryTorrasqueSummonUltralisk": {
		Name: "Torrasque",
		Text: "The Ultralisk morphs into an egg when it dies. If the egg isn't killed within 4 seconds, a new Ultralisk is born.",
	},
	"KerriganPsionicShift": {
		Name: "Psionic Shift",
		Text: "Activate to teleport to a nearby location, dealing 50 damage to nearby enemies. Generates 300% increased Assimilation Shields from the damage dealt.",
	},
	"KerriganQueensRush": {
		Name: "Queen's Rush",
		Text: "Activate to gain 30% Movement Speed for 4 seconds. Queen's Rush is also applied for free on Takedowns.",
	},
	"L90ETCMasteryDeathMetal": {
		Name: "Death Metal",
		Text: "Upon dying, a ghost uses Mosh Pit at E.T.C's location.",
	},
	"L90ETCMasteryFaceMeltLoudSpeakers": {
		Name: "Loud Speakers",
		Text: "Increases Face Melt range and knockback by 50%.",
	},
	"L90ETCMasteryFaceMeltPinballWizard": {
		Name: "Pinball Wizard",
		Text: "Face Melt does 300% more damage to enemies recently affected by Powerslide.",
	},
	"L90ETCMasteryGuitarSoloGroupies": {
		Name: "Groupies",
		Text: "Guitar Solo also heals nearby allied Heroes for 50 Health every second while it is active.",
	},
	"L90ETCMasteryMoshPitTourBus": {
		Name: "Tour Bus",
		Text: "Mosh Pit refreshes the cooldown of Powerslide. E.T.C. can Powerslide during Mosh Pit, which also increases its duration by 2 seconds.",
	},
	"L90ETCMasteryStageDiveCrowdPleaser": {
		Name: "Crowd Pleaser",
		Text: "Stage Dive's impact area is 50% bigger, and its cooldown is reduced by 10 seconds for every enemy Hero hit.",
	},
	"L90ETCMasteryStageDiveRockGod": {
		Name: "Rock God!",
		Text: "Rock God! \nReduces the cast and air time of Stage Dive by 50%, and increases the damage dealt by 50%.",
	},
	"LeoricBurningDespairTalent": {
		Name: "Burning Despair",
		Text: "Deal 40 damage per second to nearby enemies. The damage and area is increased by 100% while Drain Hope is active.",
	},
	"LeoricDrainHopeCrushingHopeTalent": {
		Name: "Crushing Hope",
		Text: "If Drain Hope lasts its full duration, it instantly deals an additional 15% of the target's maximum Health as damage. This damage does not heal Leoric.",
	},
	"LeoricDrainHopeHardenedBonesTalent": {
		Name: "Hardened Bones",
		Text: "Gain 20 Armor while Drain Hope is active. If Drain Hope lasts its full duration, retain the Armor bonus for 3.1 seconds.",
	},
	"LeoricHeroicAbilityEntomb": {
		Name: "Entomb",
		Text: "Create an unpassable tomb for 4 seconds.",
	},
	"LeoricHeroicAbilityMarchoftheBlackKing": {
		Name: "March of the Black King",
		Text: "Leoric becomes Unstoppable and walks forward, swinging his mace 3 times. Enemies hit take 250 damage, and Heroes hit heal Leoric for 12% of his maximum Health.",
	},
	"LeoricMasteryBuriedAliveEntomb": {
		Name: "Buried Alive",
		Text: "The duration of Entomb is lowered by 1 second, but enemies trapped inside are Silenced and take 50 damage every 1 second.",
	},
	"LeoricMasteryConsumeVitalitySkeletalSwing": {
		Name: "Consume Vitality",
		Text: "Skeletal Swing's cooldown is reduced by 2 seconds. Enemy Heroes hit heal Leoric for 4% of his maximum Health.",
	},
	"LeoricMasteryDeathMarchMarchoftheBlackKing": {
		Name: "Death March",
		Text: "The final swing of March of the Black King also applies Drain Hope to nearby enemy Heroes.",
	},
	"LeoricMasteryDrainMomentumDrainHope": {
		Name: "Drain Momentum",
		Text: "Drain Hope no longer reduces Leoric's Movement Speed. If Drain Hope lasts its full duration, gain 30% Movement Speed for 4 seconds.",
	},
	"LeoricMasteryFealtyUntoDeathUndying": {
		Name: "Fealty Unto Death",
		Text: "When a nearby friendly or enemy Minion dies, Leoric restores 1% of his maximum Health and 5 Mana.",
	},
	"LeoricMasteryGhastlyReachSkeletalSwing": {
		Name: "Ghastly Reach",
		Text: "Increase the range of Skeletal Swing by 25%.",
	},
	"LeoricMasteryHopelessnessDrainHope": {
		Name: "Hopelessness",
		Text: "Increase the range of Drain Hope by 25%.",
	},
	"LeoricMasteryLingeringApparitionWraithWalk": {
		Name: "Lingering Apparition",
		Text: "Increases the duration of Wraith Walk by 80%.",
	},
	"LeoricMasteryOsseinRenewal": {
		Name: "Ossein Renewal",
		Text: "Activate to heal Leoric for 30% of his maximum Health over 5 seconds. Regeneration Globes lower the cooldown of this Ability by 20 seconds.",
	},
	"LeoricMasteryParalyzingRageSkeletalSwing": {
		Name: "Paralyzing Rage",
		Text: "Increase the Slow of Skeletal Swing by 20%.",
	},
	"LeoricMasteryRoyalFocusWraithWalk": {
		Name: "Royal Focus",
		Text: "Wraith Walk increases the damage of the next Skeletal Swing within 5 seconds by 50%. Each enemy Hero hit by this Skeletal Swing reduces the cooldown of Wraith Walk by 7 seconds.",
	},
	"LeoricMasterySpectralLeech": {
		Name: "Spectral Leech",
		Text: "Basic Attacks against enemy Heroes deal bonus damage equal to 2.5% of the Hero's maximum Health and heal Leoric for twice that amount.",
	},
	"LeoricMasteryUnyieldingDespairDrainHope": {
		Name: "Unyielding Despair",
		Text: "Drain Hope's cooldown is reduced by 0.5 seconds for every 1 second it is active. If Drain Hope lasts its full duration, its cooldown is reduced by an additional 2 seconds.",
	},
	"LeoricMasteryWillingVesselDrainHope": {
		Name: "Willing Vessel",
		Text: "Increase the healing from Drain Hope to 25% of Leoric's maximum Health. If Drain Hope lasts its full duration, Leoric instantly heals for an additional 5% of his maximum Health.",
	},
	"LeoricMithrilMaceTalent": {
		Name: "Mithril Mace",
		Text: "Increase Leoric's Attack Speed by 20%.\nRepeatable Quest: Nearby enemy Minion deaths grant 0.3% permanent Attack Speed, and Takedowns grant 3% permanent Attack Speed, up to 30%.",
	},
	"LeoricShroudoftheDeadKingTalent": {
		Name: "Shroud of the Dead King",
		Text: "Activate to become Protected for 3 seconds, preventing all damage.",
	},
	"LeoricSkeletalSwingKneelPeasantsTalent": {
		Name: "Kneel, Peasants!",
		Text: "Skeletal Swing deals 100% more damage to Minions, Mercenaries, and Monsters.",
	},
	"LeoricWraithWalkOminousWraithTalent": {
		Name: "Ominous Wraith",
		Text: "Increase Wraith Walk's duration by 100%. Enemy Heroes that come in contact with the wraith deal 50% less damage for 4 seconds.",
	},
	"LiLiHeroicAbilityJugof1000Cups": {
		Name: "Jug of 1,000 Cups",
		Text: "Rapidly tosses brew to the most injured nearby allies within 10 range, restoring 1632 Health over 6 seconds.",
	},
	"LiLiHeroicAbilityWaterDragon": {
		Name: "Water Dragon",
		Text: "Li Li channels for 2 seconds, summoning a Water Dragon that hits the nearest enemy Hero within 12 range and all enemies near them, dealing 318 damage and slowing their Movement Speed by 70% for 4 seconds.",
	},
	"LiLiMasteryBlindingWindGaleForce": {
		Name: "Gale Force",
		Text: "Increases Blinding Wind damage by 40%.",
	},
	"LiLiMasteryBlindingWindHinderingWinds": {
		Name: "Hindering Winds",
		Text: "Blinding Wind also slows enemy Movement Speed by 25% for 2 seconds.",
	},
	"LiLiMasteryBlindingWindLingeringBlind": {
		Name: "Lingering Blind",
		Text: "Increases the duration of Blinding Wind by 50%.",
	},
	"LiLiMasteryBlindingWindMassVortex": {
		Name: "Mass Vortex",
		Text: "Increases the number of enemies hit by Blinding Wind from 2 to 4.",
	},
	"LiLiMasteryBlindingWindSurgingWinds": {
		Name: "Surging Winds",
		Text: "Gain 5% Spell Power for 8 seconds for every enemy hit by Blinding Wind. Additional enemies hit refresh the duration of this buff and further increase Spell Power. Stacks up to 4 times.",
	},
	"LiLiMasteryCloudSerpentBringerofGifts": {
		Name: "Bringer of Gifts",
		Text: "Cloud Serpent also heals the target for 196 Health and 0 Mana.",
	},
	"LiLiMasteryCloudSerpentLightningSerpent": {
		Name: "Lightning Serpent",
		Text: "Cloud Serpent's attacks bounce to 3 nearby enemy targets, dealing 14 damage each.",
	},
	"LiLiMasteryCloudSerpentMendingSerpent": {
		Name: "Mending Serpent",
		Text: "Cloud Serpent heals the friendly unit for 22 Health each time it attacks.",
	},
	"LiLiMasteryCloudSerpentSerpentSidekick": {
		Name: "Serpent Sidekick",
		Text: "Li Li also gains a Cloud Serpent whenever it is cast on another ally.",
	},
	"LiLiMasteryCloudSerpentTimelessCreature": {
		Name: "Timeless Creature",
		Text: "Increases the duration of Cloud Serpent by 50%.",
	},
	"LiLiMasteryFastFeetElusiveFeet": {
		Name: "Elusive Feet",
		Text: "When Fast Feet is triggered, gain 75 Physical Armor against the next 2 Hero Basic Attacks for 10 seconds, reducing the damage taken by 75%.\nThis effect has a 10 second cooldown.",
	},
	"LiLiMasteryFastFeetSafetySprint": {
		Name: "Safety Sprint",
		Text: "Increases Fast Feet's Movement Speed bonus from 10% to 20%. ",
	},
	"LiLiMasteryHealingBrewPitchPerfect": {
		Name: "Pitch Perfect",
		Text: "After casting Healing Brew, its cost is reduced by 10 Mana for 6 seconds.  This effect does not stack.",
	},
	"LiLiMasteryHealingBrewProToss": {
		Name: "Pro Toss",
		Text: "Increases the range of Healing Brew by 30%.",
	},
	"LiLiMasteryHealingBrewTheGoodStuff": {
		Name: "The Good Stuff",
		Text: "Healing Brew heals for an additional 60 Health over 6 seconds.",
	},
	"LiLiMasteryHealingBrewTwoForOne": {
		Name: "Two For One",
		Text: "Increases the number of allies healed by Healing Brew to 2, but increases the cooldown by 1.5 seconds.",
	},
	"LiLiMasteryJugof1000CupsJugof1000000Cups": {
		Name: "Jug of 1,000,000 Cups",
		Text: "Jug of 1,000 Cups hits two targets at a time.",
	},
	"LiLiMasteryKungFuHustle": {
		Name: "Kung Fu Hustle",
		Text: "Ability cooldowns refresh 150% faster while Fast Feet is active.",
	},
	"LiLiMasteryMagicalEssence": {
		Name: "Magical Essence",
		Text: "Magical Essence\nIncreases the range of Healing Brew, Cloud Serpent, and Blinding Wind by 25%.",
	},
	"LiLiMasteryShakeItOff": {
		Name: "Shake It Off",
		Text: "Being Stunned or Rooted gives 50 Armor, reducing the damage taken by 50% for 4 seconds.\nThis effect has a 10 second cooldown.",
	},
	"LiLiMasteryWaterDragonDoubleDragon": {
		Name: "Double Dragon",
		Text: "After hitting a target with Water Dragon, another Water Dragon is summoned at the point of impact.",
	},
	"LostVikingsGoGoGo64KBMarathon": {
		Name: "64 KB Marathon",
		Text: "Gain an additional 40% Movement Speed when activating Go Go Go! that decays over 4 seconds.  Additionally, the Vikings will break out of Roots and Slows.",
	},
	"LostVikingsHeroicAbilityLongboatRaid": {
		Name: "Longboat Raid!",
		Text: "Hop into an Unstoppable Longboat that fires at nearby enemies for 112 damage per second and can fire a mortar that deals 205 damage in an area.  The boat has increased Health for each Viking inside. If the boat is destroyed by enemies, all Vikings are stunned for 1.5 seconds. Lasts 15 seconds.\nRequires all surviving Vikings to be nearby.",
	},
	"LostVikingsHeroicAbilityPlayAgain": {
		Name: "Play Again!",
		Text: "Summon, fully heal, and revive all Lost Vikings at target location after a Viking channels for 2 seconds. \nOnly one Viking may attempt to summon at a time.",
	},
	"LostVikingsHeroicAbilityTheSequel": {
		Name: "The Sequel!",
		Text: "Reduces the Lost Vikings' death timer by 50%.",
	},
	"LostVikingsMasteryBaleogTheFierce": {
		Name: "Baleog the Fierce",
		Text: "Baleog's Basic Attacks increase his Attack Speed by 8%, stacking up to 5 times. After 3 seconds of not attacking, these stacks will rapidly decay.",
	},
	"LostVikingsMasteryCheckpointReached": {
		Name: "Checkpoint Reached",
		Text: "10 seconds after using Play Again!, any dead Vikings are revived and summoned again, and all are healed to full. ",
	},
	"LostVikingsMasteryErikTheSwift": {
		Name: "Erik the Swift",
		Text: "Permanently increases Erik's base Movement Speed by 10%, and as long as Erik is moving he heals 34 Health per second.",
	},
	"LostVikingsMasteryExplosiveAttacks": {
		Name: "Explosive Attacks",
		Text: "Increases Baleog's splash damage against non-Heroes to 100%.",
	},
	"LostVikingsMasteryHunkaBurningOlaf": {
		Name: "Hunka' Burning Olaf",
		Text: "Olaf deals 34 damage every second to nearby enemies.",
	},
	"LostVikingsMasteryImpatienceIsAVirtue": {
		Name: "Impatience Is a Virtue",
		Text: "Enemies damaged by a Viking's Basic Attack reduce the cooldown of all Viking Abilities by 0.2 seconds.",
	},
	"LostVikingsMasteryItsASabotage": {
		Name: "It's a Sabotage!",
		Text: "Erik's Basic Attacks against Structures light them on fire for 10 seconds, causing 135 damage and destroying 5 Ammo.",
	},
	"LostVikingsMasteryJump": {
		Name: "Jump!",
		Text: "Makes all Vikings Invulnerable and able to pass over enemies for 1.5 seconds.",
	},
	"LostVikingsMasteryLargeAndInCharge": {
		Name: "Large and In Charge",
		Text: "When Olaf charges enemies, they are stunned for 1 second.",
	},
	"LostVikingsMasteryNordicAttackSquad": {
		Name: "Nordic Attack Squad",
		Text: "Activate to have all Viking Basic Attacks deal bonus damage equal to 1% of a Hero's maximum Health for 5 seconds.",
	},
	"LostVikingsMasteryNorseForce": {
		Name: "Norse Force!",
		Text: "All Vikings gain a 110 to 220 point Shield, increasing in strength for each Viking alive. Lasts 4 seconds.",
	},
	"LostVikingsMasteryOlafTheStout": {
		Name: "Olaf the Stout",
		Text: "Every 5 seconds, Olaf gains 75 Physical Armor against the next enemy Hero Basic Attack, reducing the damage taken by 75%.\nStores up to 2 charges.",
	},
	"LostVikingsMasteryPainDontHurt": {
		Name: "Pain Don't Hurt",
		Text: "Baleog's Basic Attacks and splash damage heal for 20% of the damage dealt. Healing is doubled against Heroes.",
	},
	"LostVikingsMasterySpinToWin": {
		Name: "Spin To Win!",
		Text: "Activate to have each Viking deal 101 damage to nearby enemies.",
	},
	"LostVikingsMasterySpyGames": {
		Name: "Spy Games",
		Text: "After standing still for 3 seconds, Erik gains Stealth and his Sight Radius is increased by 75%. The Stealth persists for 3 seconds after moving.",
	},
	"LostVikingsMasteryTheSequel": {
		Name: "The Sequel!",
		Text: "Reduces the Lost Vikings' death timer by 50%.",
	},
	"LostVikingsMasteryVikingBribery": {
		Name: "Viking Bribery",
		Text: "Kill Minions to bribe a Mercenary",
	},
	"LostVikingsMasteryWereOnABoat": {
		Name: "Ragnarok 'n' Roll!",
		Text: "The Longboat can attack two targets at once and the range of its Mortar is increased by 100%.",
	},
	"LostVikingsTalentFuryoftheStorm": {
		Name: "Fury of the Storm",
		Text: "Every 5 seconds, the next Basic Attack will deal an additional 41 damage to the target, and 105 damage to all nearby Minions, Mercenaries and Monsters.  Each Viking has their own cooldown.",
	},
	"LucioAmpItUpBringItTogether": {
		Name: "Bring it Together",
		Text: "If at least 2 Allied Heroes are nearby when Amp It Up is cast, Crossfade's Healing Boost is increased by 45% for the duration.",
	},
	"LucioAmpItUpMaximumTempoQuest": {
		Name: "Maximum Tempo",
		Text: "Quest: Kill 5 enemy Heroes while Crossfade's Speed Boost is active.\nReward: Permanently increase the Amp It Up bonus for Crossfade's Speed Boost to 60%.",
	},
	"LucioAmpItUpRejuvenescencia": {
		Name: "Rejuvenescência",
		Text: "While Amp It Up is active, Crossfade's Healing Boost heals for an additional 3.2% of the target's maximum health each second.",
	},
	"LucioAmpItUpSonicAmplifier": {
		Name: "Sonic Amplifier",
		Text: "If at least 2 Allied Heroes are nearby when Amp It Up is cast, Crossfade's radius is increased by 50% for the duration.",
	},
	"LucioAmpItUpSynaesthesiaAuditiva": {
		Name: "Synaesthesia Auditiva",
		Text: "The initial cast of Amp It Up removes all Stun, Slow, and Root effects from Allies, and also removes Slow and Root effects on Lúcio.",
	},
	"LucioAmpItUpUpTheFrequency": {
		Name: "Up the Frequency",
		Text: "Amp It Up's mana cost is reduced from 100 to 80 and dealing Basic Attack damage to enemy Heroes also reduces the cooldown of Amp It Up by 0.5 seconds.",
	},
	"LucioAmptItUpBonusTrack": {
		Name: "Bonus Track",
		Text: "Swapping Crossfade tracks while Amp It Up is active will set Amp It Up back to its max duration.\nCan only happen once per cast.",
	},
	"LucioBackInTheMix": {
		Name: "Back in the Mix",
		Text: "Heal for 260 upon entering a Stun, Silence, or Time Stop effect. This cannot happen more than once every 10 seconds.",
	},
	"LucioBoombox": {
		Name: "Boombox",
		Text: "Place a boombox that plays Lúcio's active Crossfade track for 30 seconds. Its volume adjusts with Amp It Up.\nCrossfade tracks do not stack.",
	},
	"LucioCrossfadeBeatMixing": {
		Name: "Beat Mixing",
		Text: "After switching Crossfade tracks, if the next song plays for at least 1.5 seconds, Lúcio gains a 100 point Shield that lasts indefinitely.\nSwitching tracks again removes the Shield.",
	},
	"LucioCrossfadePartyMixQuest": {
		Name: "Party Mix",
		Text: "Quest: Play Lúcio's Crossfade tracks to nearby Allies for a total of 8 minutes. Multiple Allies provide additional time.\nReward: Permanently increase Crossfade's range by 20%.",
	},
	"LucioCrossfadeSpeedBoostWeMoveTogether": {
		Name: "We Move Together",
		Text: "Increase the base Movement Speed bonus of Crossfade's Speed Boost to 20%.",
	},
	"LucioReverseAmp": {
		Name: "Reverse Amp",
		Text: "Blast Lúcio's Crossfade track at enemy Heroes for 4 seconds, causing Healing Boost to inflict 112 damage per second and Speed Boost to slow for 45%.\nThis ability is unaffected by Crossfade talents.\nPassive: Increase Amp It Up's duration to 4.1 seconds.",
	},
	"LucioReverseAmpNonstopRemix": {
		Name: "Nonstop Remix",
		Text: "While Reverse Amp is affecting at least 2 enemy Heroes, its duration lasts indefinitely.",
	},
	"LucioSoundBarrier": {
		Name: "Sound Barrier",
		Text: "After 1 second, Lúcio and nearby Allied Heroes gain a 1392 point shield that rapidly decays over 6 seconds.",
	},
	"LucioSoundBarrierBossaNova": {
		Name: "Bossa Nova",
		Text: "Reduce the cooldown of Sound Barrier to 30 seconds, but the shield now decays over 4 seconds.",
	},
	"LucioSoundwaveChaseTheBassQuest": {
		Name: "Chase the Bass",
		Text: "Quest: Hit 20 enemy Heroes with Soundwave.\nReward: Increase Soundwave's arc by 50% and its range by 20%.",
	},
	"LucioSoundwaveHeliotropics": {
		Name: "Heliotropics",
		Text: "Soundwave Blinds enemies hit for 2.5 seconds.",
	},
	"LucioSoundwaveOffTheWall": {
		Name: "Off the Wall",
		Text: "If Soundwave is cast while having Wall Ride's speed bonus, the cooldown is reduced by 3 seconds.",
	},
	"LucioSoundwaveSubwoofer": {
		Name: "Subwoofer",
		Text: "Enemies within the first half of Soundwave's range are knocked back 50% farther.",
	},
	"LucioWallRideAccelerando": {
		Name: "Accelerando",
		Text: "Wall Ride's Movement Speed bonus gradually increases to 40% over 4 seconds while Lúcio maintains its effect.",
	},
	"LucioWallRideCantStopWontStop": {
		Name: "Can't Stop, Won't Stop",
		Text: "While Wall Ride's Movement Speed bonus is active, Lúcio is immune to all enemy Slow and Root effects.",
	},
	"LucioWallRideHardStyle": {
		Name: "Hard Style",
		Text: "Gain 25 Armor while Wall Ride's Movement Speed bonus is active, reducing damage taken by 25%.",
	},
	"LucioWallRideSlip": {
		Name: "Slip",
		Text: "Passing through or near an enemy Hero during Wall Ride's effect increases its Movement Speed bonus to 70% for 1 second.",
	},
	"MalfurionCombatStyleElunesGrace": {
		Name: "Elune's Grace",
		Text: "Increases the range of Regrowth, Moonfire, and Entangling Roots by 25%.",
	},
	"MalfurionCombatStyleShandosClarity": {
		Name: "Shan'do's Clarity",
		Text: "Reduces Innervate's cooldown by 8 seconds.",
	},
	"MalfurionHeroicAbilityTranquility": {
		Name: "Tranquility",
		Text: "Heals 80 Health per second to nearby allied Heroes over 10 seconds.",
	},
	"MalfurionHeroicAbilityTwilightDream": {
		Name: "Twilight Dream",
		Text: "After 0.5 seconds, deal 374 damage in a large area around Malfurion, Silencing enemies making them unable to use Abilities for 3 seconds.",
	},
	"MalfurionMasteryAstralCommunion": {
		Name: "Astral Communion",
		Text: "Activate to Channel for 1 second, and then instantly teleport and cast Twilight Dream at the targeted location.  Cooldown is shared with Twilight Dream.\nPassive: Increase Twilight Dream's Silence duration to 4 seconds.",
	},
	"MalfurionMasteryFullMoonfire": {
		Name: "Full Moonfire",
		Text: "Increases Moonfire's radius by 50% and reduces its Mana cost by 5.",
	},
	"MalfurionMasteryHinderingMoonfire": {
		Name: "Hindering Moonfire",
		Text: "Moonfire Slows enemy Heroes by 25% for 2 seconds.",
	},
	"MalfurionMasteryLifeSeed": {
		Name: "Life Seed",
		Text: "Nearby damaged allied Heroes will automatically gain the heal over time portion of Regrowth. This effect has a 20 second cooldown.",
	},
	"MalfurionMasteryLunarShower": {
		Name: "Lunar Shower",
		Text: "Using Moonfire reduces the cooldown of the next Moonfire by 0.5 seconds, and increases the damage by 15%. Stacks up to 3 times and resets after 6 seconds.",
	},
	"MalfurionMasteryMoonburn": {
		Name: "Moonburn",
		Text: "Increase Moonfire's damage to Minions by 75%.\nRepeatable Quest: Every 5 Minions killed nearby increase this bonus by 1%, up to 25%. After reaching a 25% bonus, Moonburn's bonus damage also applies to Mercenaries and Monsters.",
	},
	"MalfurionMasterySerenity": {
		Name: "Serenity",
		Text: "Increases Tranquility's healing by 25% and it also restores 5 Mana per second.",
	},
	"MalfurionMasteryStranglingVinesEntanglingRoots": {
		Name: "Strangling Vines",
		Text: "Enemy Heroes Rooted by Entangling Roots receive 30% less healing from all sources.",
	},
	"MalfurionMasteryTenaciousRootsEntanglingRoots": {
		Name: "Tenacious Roots",
		Text: "Entangling Roots grows 25% larger, lasts 25% longer, and Roots targets for 0.2 seconds longer.",
	},
	"MalfurionMasteryVengefulRoots": {
		Name: "Vengeful Roots",
		Text: "Entangling Roots spawns a Treant that deals 70 damage per second and lasts 10 seconds.\nRepeatable Quest: Hitting enemy Heroes with Entangling Roots permanently increases the Treant's damage per second by 8.",
	},
	"MalfurionRevitalizeInnervateTalent": {
		Name: "Revitalize",
		Text: "Using Innervate also grants Malfurion 50 Mana and causes his Basic Ability cooldowns to refresh 50% faster for 5 seconds.",
	},
	"MalfurionYserasGift": {
		Name: "Ysera's Gift",
		Text: "While above 75% Health, Regrowth's additional healing-over-time is increased by 40%.",
	},
	"MalthaelAngelOfDeath": {
		Name: "Angel of Death",
		Text: "Last Rites heals for 100% of the damage dealt, and its current and future cooldown reduction bonuses are doubled.",
	},
	"MalthaelBlackHarvest": {
		Name: "Black Harvest",
		Text: "Quest: Apply Reaper's Mark to Heroes for a total of 150 seconds.\nReward: Permanently increase Reaper's Mark's duration by 2 seconds.",
	},
	"MalthaelColdHand": {
		Name: "Cold Hand",
		Text: "Soul Rip Slows enemies by 20% for 2.5 seconds.",
	},
	"MalthaelDeathsReach": {
		Name: "Death's Reach",
		Text: "Increase Wraith Strike's range by 35%.",
	},
	"MalthaelDieAlone": {
		Name: "Die Alone",
		Text: "Soul Rip deals 75% more damage if it hits only one Hero and no other enemies.",
	},
	"MalthaelEtherealExistence": {
		Name: "Ethereal Existence",
		Text: "Gain 10 Physical Armor per enemy Hero afflicted by Reaper's Mark, up to a maximum of 40.",
	},
	"MalthaelFearTheReaper": {
		Name: "Fear the Reaper",
		Text: "Activate to increase Movement Speed by 25% and pass through other units for 4 seconds.",
	},
	"MalthaelFinalCurtain": {
		Name: "Final Curtain",
		Text: "Death Shroud leaves a trail in its wake for 4 seconds, applying Reaper's Mark to enemies in its area.",
	},
	"MalthaelInevitableEnd": {
		Name: "Inevitable End",
		Text: "Activate to become Unstoppable for 2 seconds, but remove all active Reaper's Marks.",
	},
	"MalthaelLastRites": {
		Name: "Last Rites",
		Text: "Apply a death sentence to an enemy Hero that, after 2 seconds, deals damage equal to 50% of their missing Health.\nRepeatable Quest: Enemies killed while under the effect of Last Rites permanently reduce its cooldown by 5 seconds, to a minimum of 15 seconds.",
	},
	"MalthaelMassacre": {
		Name: "Massacre",
		Text: "Wraith Strike now damages and applies Reaper's Mark to enemies around its target.",
	},
	"MalthaelMementoMori": {
		Name: "Memento Mori",
		Text: "Reaper's Mark deals 80% increased damage after afflicting an enemy for more than 4 seconds.",
	},
	"MalthaelMortality": {
		Name: "Mortality",
		Text: "When damaging a Hero, Wraith Strike deals bonus damage equal to 4% of the Hero's maximum Health.",
	},
	"MalthaelNoOneCanStopDeath": {
		Name: "No One Can Stop Death",
		Text: "Activate while dead to immediately respawn at the Altar but increase Malthael's next respawn time by 25%.",
	},
	"MalthaelOnAPaleHorse": {
		Name: "On a Pale Horse",
		Text: "Gain an additional 20% Movement Speed while mounted.",
	},
	"MalthaelReaperOfSouls": {
		Name: "Reaper of Souls",
		Text: "Increase Tormented Souls' duration by 1 second.  While Tormented Souls is active, Hero Takedowns extend its duration by 4 seconds.",
	},
	"MalthaelShroudOfWisdom": {
		Name: "Shroud of Wisdom",
		Text: "After 2 seconds, gain 50 Spell Armor for 4 seconds.",
	},
	"MalthaelSoulCollector": {
		Name: "Soul Collector",
		Text: "Reduce Soul Rip's cooldown by 0.5 seconds and increase its range by 50%.",
	},
	"MalthaelSoulSiphon": {
		Name: "Soul Siphon",
		Text: "Increase Soul Rip's bonus healing from Heroes to 4% of the Hero's maximum Health.",
	},
	"MalthaelThrowingShade": {
		Name: "Throwing Shade",
		Text: "Quest: Hit 20 Heroes with Death Shroud.\nReward: Permanently increase Death Shroud's range by 33.3%, reduce its cooldown by 4 seconds, and reduce its Mana cost from 50 to 25.",
	},
	"MalthaelTormentedSouls": {
		Name: "Tormented Souls",
		Text: "Gain 10 Armor and unleash a torrent of souls, continually applying Reaper's Mark to nearby enemies for 4 seconds.",
	},
	"MalthaelTouchOfDeath": {
		Name: "Touch of Death",
		Text: "Activate to reduce healing received by Heroes afflicted by Reaper's Mark by 50% for 4 seconds.",
	},
	"MedicBlastShield": {
		Name: "Blast Shield",
		Text: "Heroes hit by Displacement Grenade generate 2 Energy and grant Lt. Morales a Shield equal to 6% of her maximum Health, stacking up to 5 times.",
	},
	"MedicCaduceusFeedback": {
		Name: "Caduceus Feedback",
		Text: "Increases Attack Range by 1.1 and Basic Attacks against Heroes generate 4 Energy.",
	},
	"MedicCaduceusReactor2dot0": {
		Name: "Caduceus Reactor 2.0",
		Text: "Increase the healing provided by Caduceus Reactor from 2% to 6% of Lt. Morales's maximum Health per second.",
	},
	"MedicCellularReactor": {
		Name: "Cellular Reactor",
		Text: "Consume 30 Energy to heal Lt. Morales for 40% of her maximum Health over 4 seconds. Caduceus Reactor is disabled while this is active.",
	},
	"MedicClear": {
		Name: "Clear!",
		Text: "Quest: Hit Heroes with Displacement Grenade.\nReward: After hitting 15 Heroes, reduce the cooldown of Displacement Grenade from 12 seconds to 9 seconds.\nReward: After hitting 30 Heroes, increase the detonation area of Displacement Grenade by 25%.",
	},
	"MedicEMPGrenade": {
		Name: "EMP Grenade",
		Text: "Displacement Grenade deals an additional +Inf damage over 2 seconds, and up to 400 bonus damage to Shields.",
	},
	"MedicExtendedCare": {
		Name: "Extended Care",
		Text: "Increase the range of Healing Beam by 40%.",
	},
	"MedicFirstResponder": {
		Name: "First Responder",
		Text: "While above 70 Energy, Healing Beam's healing is increased by 25%.",
	},
	"MedicHeroicAbilityMedivacDropship": {
		Name: "Medivac Dropship",
		Text: "Target a location for a Medivac transport. For up to 10.5 seconds before takeoff, allies can right-click to enter the Medivac.",
	},
	"MedicHeroicAbilityStimDrone": {
		Name: "Stim Drone",
		Text: "Grant an allied Hero 75% Attack Speed and 25% Movement Speed for 10 seconds.",
	},
	"MedicHyperactivity": {
		Name: "Hyperactivity",
		Text: "Reduce the cooldown of Stim Drone from 90 seconds to 45 seconds, and increase the Movement Speed bonus from 25% to 50%.",
	},
	"MedicLifeSupport": {
		Name: "Life Support",
		Text: "Generate 2 Energy each time Safeguard reduces damage, up to a maximum of 20 Energy per use.",
	},
	"MedicPhysicalTherapy": {
		Name: "Physical Therapy",
		Text: "Safeguard removes all Slows from its target. When a Slow is removed this way, reduce the cooldown of Safeguard by 4 seconds.",
	},
	"MedicProlongedSafeguard": {
		Name: "Prolonged Safeguard",
		Text: "Increase Safeguard's duration by 50%.",
	},
	"MedicReinforcements": {
		Name: "Reinforcements",
		Text: "Activate to call down a Medivac Dropship at your Altar for your allies. After 10 seconds, or when the Ability is activated again, the Medivac will travel to the original cast location to unload its passengers. Cooldown is shared with Medivac Dropship.\nPassive: Reduce the cooldown of Medivac Dropship from 60 to 30 seconds.\nDoes not require the Medivac Dropship Talent.",
	},
	"MedicSafeZone": {
		Name: "Safe Zone",
		Text: "After exiting the Medivac, allies are Protected from all damage for 2 seconds.",
	},
	"MedicSecondOpinion": {
		Name: "Second Opinion",
		Text: "Hitting 2 or more Heroes with Displacement Grenade reduces its cooldown to 2 seconds.",
	},
	"MedicShieldSequencer": {
		Name: "Shield Sequencer",
		Text: "Safeguard gains a second charge. These charges cannot stack on a single target.",
	},
	"MedicSystemShock": {
		Name: "System Shock",
		Text: "Heroes hit by Displacement Grenade deal 30% less damage for 4 seconds.",
	},
	"MedicTraumaTrigger": {
		Name: "Trauma Trigger",
		Text: "Taking damage while below 40% Health grants 30 Armor for 3 seconds.\nThis effect has a 60 second cooldown.",
	},
	"MedicVanadiumPlating": {
		Name: "Vanadium Plating",
		Text: "While an ally affected by Safeguard is Stunned, Safeguard grants an additional 25 Armor and its duration is paused.",
	},
	"MedivhArcaneBrilliance": {
		Name: "Arcane Brilliance",
		Text: "Activate to instantly restore 200 mana to all nearby allied Heroes and grant them 10% Spell Power for 10 seconds.",
	},
	"MedivhArcaneRiftArcaneCharge": {
		Name: "Arcane Charge",
		Text: "When Arcane Rift damages an enemy Hero, the next Arcane Rift deals 30% more damage.",
	},
	"MedivhArcaneRiftGuardianOfTirisfal": {
		Name: "Guardian of Tirisfal",
		Text: "Minions and Catapults hit by Arcane Rift are instantly killed.",
	},
	"MedivhArcaneRiftManaAdept": {
		Name: "Mana Adept",
		Text: "Increases Arcane Rift's Mana refund for hitting enemy Heroes by 60%.",
	},
	"MedivhArcaneRiftTheMastersTouch": {
		Name: "The Master's Touch",
		Text: "Quest: Hit 30 enemy Heroes with Arcane Rift without dying.\nReward: Permanently increases the damage dealt by 75 and cooldown reduction for hitting a Hero by 1 second.",
	},
	"MedivhDustOfAppearance": {
		Name: "Dust of Appearance",
		Text: "Activate to reveal enemy Heroes in the area surrounding Medivh for 10 seconds. Does not cancel Raven Form.",
	},
	"MedivhForceOfWillArcaneExplosion": {
		Name: "Arcane Explosion",
		Text: "The first time Force of Will absorbs damage, it deals 183 damage to nearby enemies.",
	},
	"MedivhForceOfWillCircleOfProtection": {
		Name: "Circle of Protection",
		Text: "An untalented Force of Will is also applied to allies near the target.",
	},
	"MedivhForceOfWillEnduringWill": {
		Name: "Enduring Will",
		Text: "Increases Force of Will duration by 33.3%.",
	},
	"MedivhForceOfWillReabsorption": {
		Name: "Reabsorption",
		Text: "Upon expiration, Force of Will heals the target for 50% of the damage it absorbed.",
	},
	"MedivhInvisibility": {
		Name: "Invisibility",
		Text: "Stealth an allied Hero for 20 seconds.\nStores up to 2 charges.",
	},
	"MedivhLeyLineSeal": {
		Name: "Ley Line Seal",
		Text: "After 0.5 seconds, unleash a wave of energy that places enemy Heroes in Time Stop for 3 seconds.",
	},
	"MedivhLeyLineSealMedivhCheats": {
		Name: "Medivh Cheats!",
		Text: "Increase the duration of the Time Stop by 33.3% and gain the ability to redirect the wave once while it's active.",
	},
	"MedivhPolyBomb": {
		Name: "Poly Bomb",
		Text: "Polymorph an enemy Hero for 2 seconds, Silencing them and making them unable to attack. On expiration, Poly Bomb spreads to other nearby enemy Heroes.",
	},
	"MedivhPolyBombGlyphOfPolyBomb": {
		Name: "Glyph Of Poly Bomb",
		Text: "Poly Bomb explodes 1 second sooner and the spread radius is increased by 25%.",
	},
	"MedivhPortalAstralProjection": {
		Name: "Astral Projection",
		Text: "Increases Portal cast range by 50%.",
	},
	"MedivhPortalMageArmor": {
		Name: "Mage Armor",
		Text: "The first time an Ally uses a Portal, they gain 50 Physical Armor against the next 3 enemy Heroic Basic Attacks for 5 seconds, reducing the damage taken by 50%.",
	},
	"MedivhPortalPortalMastery": {
		Name: "Portal Mastery",
		Text: "Medivh can manually place both Portal locations. \nActivate Medivh's Trait to cancel an unlinked Portal.",
	},
	"MedivhPortalQuickening": {
		Name: "Quickening",
		Text: "Reduces Portal cooldown by 50%.",
	},
	"MedivhPortalRavenFamiliar": {
		Name: "Raven Familiar",
		Text: "A Raven Familiar joins allies that use a Portal. The Raven will dive at the ally's next Basic Attack against a Hero within 5 seconds for 114 damage.",
	},
	"MedivhPortalStablePortal": {
		Name: "Stable Portal",
		Text: "Increases Portal duration by 50%.",
	},
	"MedivhTransformRavenBirdsEyeView": {
		Name: "Bird's Eye View",
		Text: "Permanently increase Raven Form's vision range by 25%. Activate to increase this bonus to 200% for 5 seconds.",
	},
	"MedivhTransformRavenRavensIntellect": {
		Name: "Raven’s Intellect",
		Text: "Raven Form increases Medivh's Mana Regeneration by 100%.",
	},
	"MedivhTransformRavenWindsOfCelerity": {
		Name: "Winds of Celerity",
		Text: "Increases Raven Form's Movement Speed bonus by 50%.",
	},
	"MonkAirAlly": {
		Name: "Air Ally",
		Text: "Place an Air Ally that grants vision and reveals a large area around it. Has 75 Health and lasts 40 seconds. Kharazim can Radiant Dash to Air Allies.\nStores up to 2 charges.",
	},
	"MonkBlazingFistsDeadlyReach": {
		Name: "Blazing Fists",
		Text: "Every 3rd Basic Attack reduces the cooldown of Deadly Reach by 0.8 seconds.",
	},
	"MonkBlindingSpeedRadiantDash": {
		Name: "Blinding Speed",
		Text: "Decreases Radiant Dash's cooldown by 2 seconds and increases the maximum number of charges by 1.",
	},
	"MonkCleansingTouchRadiantDash": {
		Name: "Cleansing Touch",
		Text: "Radiant Dashing to an ally makes them Unstoppable for 1 second.",
	},
	"MonkDashofLight": {
		Name: "Dash of Light",
		Text: "Radiant Dashing to an ally increases the healing they receive from Breath of Heaven by 75% for 3 seconds.",
	},
	"MonkEarthAlly": {
		Name: "Earth Ally",
		Text: "Place an Earth Ally that grants nearby allied Heroes 50 Physical Armor against Heroic Basic Attacks, reducing damage taken by 50%. Has 400 Health and lasts 10 seconds.  Kharazim can Radiant Dash to Earth Allies.",
	},
	"MonkEchoofHeavenBreathofHeaven": {
		Name: "Echo of Heaven",
		Text: "Breath of Heaven heals 75% of its normal amount, but heals a second time 3 seconds later.",
	},
	"MonkElevenSidedStrikeSevenSidedStrike": {
		Name: "Transgression",
		Text: "Seven-Sided Strike hits 4 additional times.",
	},
	"MonkEpiphany": {
		Name: "Epiphany",
		Text: "Activate to restore 32.8% of Kharazim's maximum Mana and refill 2 charges of Radiant Dash. ",
	},
	"MonkFistsofFuryDeadlyReach": {
		Name: "Fists of Fury",
		Text: "Increases Deadly Reach's duration by 100%.",
	},
	"MonkHeavenlyZealBreathofHeaven": {
		Name: "Heavenly Zeal",
		Text: "Increases Breath of Heaven's Movement Speed bonus to 30%.",
	},
	"MonkHeroicAbilityDivinePalm": {
		Name: "Divine Palm",
		Text: "Protect an allied Hero from death, causing them to be healed for 1141 if they take fatal damage in the next 3 seconds.",
	},
	"MonkHeroicAbilitySevenSidedStrike": {
		Name: "Seven-Sided Strike",
		Text: "Become Invulnerable and strike 7 times over 1.8 seconds. Each strike hits the highest Health nearby Hero for 7% of their maximum Health.",
	},
	"MonkInsight": {
		Name: "Insight",
		Text: "Quest: Every 3rd Basic Attack restores 14 Mana and grants a stack of Insight and gives 25% increased Move Speed for 2 seconds.\nReward: Upon getting 100 stacks of Insight, every 3rd attack also reduces the cooldown of all of Kharazim's Basic Abilities by 1.8 seconds.",
	},
	"MonkIronFists": {
		Name: "Iron Fists",
		Text: "Every 3rd Basic Attack deals 110% bonus damage and gives 25% increased Move Speed for 2 seconds.",
	},
	"MonkPeacefulReposeDivinePalm": {
		Name: "Peaceful Repose",
		Text: "Divine Palm's cooldown is set to 5 seconds if the Hero does not die.",
	},
	"MonkQuicksilverRadiantDash": {
		Name: "Quicksilver",
		Text: "Radiant Dashing to an ally gives Kharazim and the target 30% bonus Movement Speed for 3 seconds.",
	},
	"MonkSixthSense": {
		Name: "Sixth Sense",
		Text: "While Stunned or Rooted, gain 75 Physical Armor against Hero Basic Attacks for 4 seconds, reducing their damage by 75%. ",
	},
	"MonkSpiritAlly": {
		Name: "Spirit Ally",
		Text: "Place a Spirit Ally that heals allies in a large area around it for 1.9% of their maximum Health every second. Has 150 Health and lasts 10 seconds. Kharazim can Radiant Dash to Spirit Allies.",
	},
	"MonkTranscendence": {
		Name: "Transcendence",
		Text: "Every 3rd Basic Attack heals the lowest nearby allied Hero for 103 and gives 25% increased Move Speed for 2 seconds.",
	},
	"MonkWayoftheHundredFistsRadiantDash": {
		Name: "Way of the Hundred Fists",
		Text: "Radiant Dashing to an enemy launches a rapid volley of 6 Basic Attacks dealing 50% damage each.",
	},
	"MuradinBronzebeardRage": {
		Name: "Bronzebeard Rage",
		Text: "Deal 19 damage per second to nearby enemies. Basic Attacks against a Hero that is Rooted, Stunned or Slowed increase this damage by 100% for 3 seconds.",
	},
	"MuradinCombatStyleSkullcracker": {
		Name: "Skullcracker",
		Text: "Every 3rd Basic Attack against the same enemy deals 70% bonus damage and Stuns them for 0.2 seconds.",
	},
	"MuradinCombatStyleThirdWind": {
		Name: "Third Wind",
		Text: "Increases Health Restoration rate to 83 per second, and raises Health threshold to 60% Health for improved 165.9 per second Restoration.",
	},
	"MuradinGiveEmTheAxeExecutioner60DamageBonus": {
		Name: "Give 'em the Axe!",
		Text: "Attacking a Hero that is slowed, rooted, or stunned increases Muradin's Basic Attack damage by 50% for 3 seconds.",
	},
	"MuradinHeroicAbilityAvatar": {
		Name: "Avatar",
		Text: "Transform for 20 seconds, gaining 1053 Health.",
	},
	"MuradinHeroicAbilityHaymaker": {
		Name: "Haymaker",
		Text: "Stun target enemy Hero, and wind up a punch dealing 638 damage and knocking the target back, hitting enemies in the way for 638 damage and knocking them aside.",
	},
	"MuradinMasteryAvatarUnstoppableForce": {
		Name: "Unstoppable Force",
		Text: "While active, Avatar grants 20 Armor and causes Muradin's Basic Attacks to reduce the cooldowns of Thunder Clap and Dwarf Toss by 0.8 seconds.",
	},
	"MuradinMasteryDwarfTossDwarfLaunch": {
		Name: "Dwarf Launch",
		Text: "Increase the range of Dwarf Toss by 40%. Hitting an enemy Hero with Dwarf Toss reduces its cooldown by 3 seconds.",
	},
	"MuradinMasteryDwarfTossHeavyImpact": {
		Name: "Heavy Impact",
		Text: "Enemies hit by Dwarf Toss are slowed by 80% for 0.8 seconds.",
	},
	"MuradinMasteryDwarfTossLandingMomentum": {
		Name: "Landing Momentum",
		Text: "[PH] Increases the range of Dwarf Toss by 40% and increases Muradin's Movement Speed by 25% for 5 seconds upon landing with Dwarf Toss.",
	},
	"MuradinMasteryHaymakerGrandSlam": {
		Name: "Grand Slam",
		Text: "Haymaker gains a 2nd charge and its damage its increased by 25%. If a Hero dies within 3 seconds of being hit by Haymaker, instantly gain 1 charge.",
	},
	"MuradinMasteryPassiveStoneform": {
		Name: "Stoneform",
		Text: "Activate to heal Muradin for 30% of his maximum Health over 10 seconds. Second Wind is disabled during this time.",
	},
	"MuradinMasteryStormhammerSledgehammer": {
		Name: "Sledgehammer",
		Text: "Deals 350% damage to non-Heroic enemies. Destroys 4 ammo from Structures.",
	},
	"MuradinMasteryThunderburn": {
		Name: "Thunder Burn",
		Text: "Hitting an enemy Hero with Thunder Clap triggers a second explosion 2 seconds later in the same location that deals 50% damage.",
	},
	"MuradinMasteryThunderclapCrowdControl": {
		Name: "Crowd Control",
		Text: "Each enemy hit by Thunder Clap reduces its cooldown by 0.8 seconds and restores 5 Mana, to a maximum of 9 targets.",
	},
	"MuradinMasteryThunderclapHealingStatic": {
		Name: "Healing Static",
		Text: "Muradin heals for 5% of his Max Health for each Hero hit by Thunder Clap.",
	},
	"MuradinMasteryThunderclapReverberation": {
		Name: "Reverberation",
		Text: "Increases the Attack Speed slow from 25% to 50% and the duration from 2.5 seconds to 3.5 seconds.",
	},
	"MuradinMasteryThunderclapThunderstrike": {
		Name: "Thunder Strike",
		Text: "Thunder Clap deals 300% damage if only one target is hit.",
	},
	"MuradinStormboltPerfectStorm": {
		Name: "Perfect Storm",
		Text: "Quest: Every time Storm Bolt hits a Hero, increase its damage by 5.",
	},
	"MuradinStormboltPiercingBolt": {
		Name: "Piercing Bolt",
		Text: "Penetrates through the first target hit, hitting 1 additional target.",
	},
	"MurkyAFishyDeal": {
		Name: "A Fishy Deal",
		Text: "Killing a Minion with Pufferfish grants a stack of Bribe. Use 8 stacks to bribe a Mercenary, instantly defeating them. Does not work on Bosses.\nMaximum of 32 stacks of Bribe.",
	},
	"MurkyBigTunaKahuna": {
		Name: "Big Tuna Kahuna",
		Text: "Murky's maximum Health and Egg respawn time are doubled.",
	},
	"MurkyBlackLagoon": {
		Name: "Black Lagoon",
		Text: "Increase Slime's radius by 30%.",
	},
	"MurkyEggHunt": {
		Name: "Egg Hunt",
		Text: "Activate to place a fake Egg. If the fake Egg dies, it casts an untalented Slime. Maximum 3 fake Eggs.\nPassive: Spawning from his Egg grants Murky Stealth for 5 seconds.",
	},
	"MurkyEggShell": {
		Name: "Egg Shell",
		Text: "Spawning from his Egg grants Murky a Shield equal to 100% of his maximum Health. The shield lasts indefinitely.",
	},
	"MurkyFishEye": {
		Name: "Fish Eye",
		Text: "Egg's Health is increased by 100%, its sight radius is increased by 300% and it can see Stealthed enemies. \nPassive: Spawning from his Egg increases Murky's mount speed to 44.9% for 5 seconds.",
	},
	"MurkyFishOil": {
		Name: "Fish Oil",
		Text: "The Pufferfish casts Slime at its location upon landing.",
	},
	"MurkyFishTank": {
		Name: "Fish Tank",
		Text: "Basic Attacks against Heroes grant Murky 75 Physical Armor against Heroes for 3 seconds.",
	},
	"MurkyHeroicAbilityMarchoftheMurlocs": {
		Name: "March of the Murlocs",
		Text: "After 0.8 seconds, Murky commands a legion of Murlocs to march in a target direction, each one leaping onto the first enemy Hero or Structure they find. Each Murloc deals 125 damage and slows its target by 15% for 5 seconds. Murlocs deal 50% damage to Structures.",
	},
	"MurkyHeroicAbilityOctoGrab": {
		Name: "Octo-Grab",
		Text: "Murky summons an octopus to stun target enemy Hero for 3.1 seconds while he hits them for 1 damage a second.",
	},
	"MurkyLivingtheDream": {
		Name: "Living the Dream",
		Text: "Quest: Every 15 seconds Murky is alive, he gains 5% Spell Power, up to 15%. This bonus is reset upon his death.",
	},
	"MurkyMakingInky": {
		Name: "Making Inky",
		Text: "Reduces the Cooldown of Slime from 4 seconds to 2 seconds.",
	},
	"MurkyMasteryAndASharkToo": {
		Name: "... And a Shark Too!",
		Text: "Increase the damage of Octo-Grab by 13700%.",
	},
	"MurkyMasteryNeverEndingMurlocs": {
		Name: "Never-Ending Murlocs",
		Text: "March of the Murlocs can be channeled, sending little Murlocs indefinitely.",
	},
	"MurkyMasteryRejuvenatingBubble": {
		Name: "Rejuvenating Bubble",
		Text: "Safety Bubble restores 40% of Murky's maximum Health.",
	},
	"MurkyMasteryTufferfish": {
		Name: "Tufferfish",
		Text: "Pufferfish gains 50 Spell Armor and deals 50% more damage to Slimed targets.",
	},
	"MurkyMasteryWrathOfCod": {
		Name: "Wrath of Cod",
		Text: "Heroes hit by Pufferfish take additional damage equal to 15% of their maximum Health over 5 seconds.",
	},
	"MurkySlimeTime": {
		Name: "Slime Time",
		Text: "Quest: Slime enemy Heroes that are already Slimed.\nReward: After Sliming 15 Heroes, increase Slime's bonus damage by 125.\nReward: After Sliming 30 Heroes, increase Slime's slow amount to 30%.",
	},
	"MurkySlipperyWhenWet": {
		Name: "Slippery When Wet",
		Text: "Gain 50% Movement Speed and move through units while in Safety Bubble. Additionally, Safety Bubble's cooldown is reduced by 2 seconds.",
	},
	"MurkyTimetoKrill": {
		Name: "Time to Krill",
		Text: "Basic Attacks against Heroes deal an additional 7 damage a second and slow the target by 7% for 4 seconds. This effect can stack up to 5 times.",
	},
	"MurkyToxicBuildup": {
		Name: "Toxic Buildup",
		Text: "Every 3rd consecutive Basic Attack against an enemy Hero causes a free Slime to be cast upon them.",
	},
	"NecromancerHeroicAbilityPoisonNova": {
		Name: "Poison Nova",
		Text: "After a short delay, release poisonous missiles that deal 570 damage to all enemies hit over 10 seconds.",
	},
	"NecromancerHeroicAbilitySkeletalMages": {
		Name: "Skeletal Mages",
		Text: "Vector Targeting\nSummon 4 Frost Mages in a line that attack nearby enemies for 46.9 damage a second and Slow them by 30% for 2 seconds. Last up to 15 seconds.",
	},
	"NecromancerTalentAmplifyDamage": {
		Name: "Amplify Damage",
		Text: "Enemies Rooted by Bone Prison lose 25 Armor for 1.8 seconds.",
	},
	"NecromancerTalentAndarielsVisage": {
		Name: "Andariel's Visage",
		Text: "Poison Nova heals for 50% of the damage dealt.\nPassive: Skeletal Warriors deal 50% increased damage.",
	},
	"NecromancerTalentBoneArmorBacklash": {
		Name: "Backlash",
		Text: "When Bone Armor expires, nearby enemy Heroes take damage equal to 12% of their maximum Health.",
	},
	"NecromancerTalentBoneArmorShackler": {
		Name: "Shackler",
		Text: "While Bone Armor is active, nearby enemies are Slowed by 35% for 1 second. Basic Attacks against Heroes that are Stunned, Rooted, or Slowed reduce the cooldown of Bone Armor by 2 seconds.",
	},
	"NecromancerTalentBoneArmorShade": {
		Name: "Shade",
		Text: "While Bone Armor is active, Xul evades all Basic Attacks, but increases the cooldown of Bone Armor by 20 seconds.",
	},
	"NecromancerTalentBoneSpear": {
		Name: "Bone Spear",
		Text: "Deal 230 damage to enemies in a line.",
	},
	"NecromancerTalentColdHandOfDeath": {
		Name: "Cold Hand of Death",
		Text: "Increase the Slow of Frost Mages by 20%.\nPassive: Skeletal Warrior attacks Slow enemies by 30% for 2 seconds.",
	},
	"NecromancerTalentCorpseExplosion": {
		Name: "Corpse Explosion",
		Text: "When Skeletal Warriors die they explode for 55 damage in a small area. Deals 100% increased damage to non-Heroic enemies.",
	},
	"NecromancerTalentEchoesOfDeath": {
		Name: "Echoes of Death",
		Text: "Spectral Scythe spawns 2 additional scythes next to the first after 1.5 seconds, dealing 60% damage.",
	},
	"NecromancerTalentGrimScythe": {
		Name: "Grim Scythe",
		Text: "Each enemy hit by Cursed Strikes reduces its cooldown by 0.8 seconds, up to 12 seconds.",
	},
	"NecromancerTalentHarvestVitality": {
		Name: "Harvest Vitality",
		Text: "Cursed Strikes heals for 60% of the damage dealt to Heroes.",
	},
	"NecromancerTalentJailors": {
		Name: "Jailors",
		Text: "Bone Prison spawns 2 Skeletal Warriors.  These do not count towards Xul's Raise Skeleton maximum.\nQuest: Spawn 80 Skeletal Warriors.\nReward: While fixating on an enemy, Skeletal Warriors gain 50% Movement and Attack Speed.",
	},
	"NecromancerTalentKalansEdict": {
		Name: "Kalan's Edict",
		Text: "Skeletal Warrior attacks reduce the cooldown of Xul's Heroic Ability by 1%.",
	},
	"NecromancerTalentMortalWound": {
		Name: "Mortal Wound",
		Text: "Enemies hit by Spectral Scythe receive 75% less healing for 4 seconds.",
	},
	"NecromancerTalentRapidHarvest": {
		Name: "Rapid Harvest",
		Text: "Gain 5% Attack Speed for 3 seconds each time Cursed Strikes hits an enemy, up to 75%.",
	},
	"NecromancerTalentReapersToll": {
		Name: "Reaper's Toll",
		Text: "Quest: Hit 20 Heroes with Spectral Scythe.\nReward: Reduce the cooldown of Spectral Scythe by 3 seconds.",
	},
	"NecromancerTalentTragOulsEssence": {
		Name: "Trag'Oul's Essence",
		Text: "Every time a Skeletal Warrior attacks an enemy, restore 0.5% of Xul's maximum Health and 0.5% of his maximum Mana.",
	},
	"NecromancerTalentWeaken": {
		Name: "Weaken",
		Text: "Increase the Attack Speed reduction of Cursed Strikes by 35% and the duration by 1 second.",
	},
	"NovaAdvancedCloaking": {
		Name: "Advanced Cloaking",
		Text: "While Stealthed, Nova gains 2 Mana per second and 25% increased movement speed.",
	},
	"NovaCombatStyleAntiArmorShells": {
		Name: "Anti-Armor Shells",
		Text: "Nova's Basic Attacks deal 250% damage and decrease the Physical Armor of Heroic targets by 10 for 3 seconds, but her Attack Speed is proportionally slower.",
	},
	"NovaCombatStyleOneintheChamber": {
		Name: "One in the Chamber",
		Text: "After using an Ability, Nova's next Basic Attack within 3 seconds deals 80% additional damage.",
	},
	"NovaCovertMission": {
		Name: "Covert Mission",
		Text: "Enemy Minions killed near Nova grant a stack of Bribe. Hero Takedowns grant 10 stacks of Bribe. Use 25 stacks to bribe a Mercenary, instantly defeating them.  Does not work on Bosses.  Maximum of 100 stacks. If a camp is defeated entirely with Bribe, the camp respawns 50% faster.",
	},
	"NovaCovertOps": {
		Name: "Covert Ops",
		Text: "After being Stealthed for 5 seconds, Pinning Shot's slow is increased to 55% and costs no Mana. Bonus is lost after losing Stealth for 1 second.",
	},
	"NovaGhostProtocol": {
		Name: "Ghost Protocol",
		Text: "Activate to gain Stealth and become Unrevealable for 2 seconds. Nova's attacks and Abilities will also not break Stealth during this time.",
	},
	"NovaHeroicAbilityPrecisionStrike": {
		Name: "Precision Strike",
		Text: "After a 1.5 second delay, deals 456 damage to enemies within an area. Unlimited range.",
	},
	"NovaHeroicAbilityTripleTap": {
		Name: "Triple Tap",
		Text: "Locks in on the target Hero, then fires 3 shots that hit the first Hero or Structure they come in contact with for 338 damage each.",
	},
	"NovaHoloStability": {
		Name: "Holo Stability",
		Text: "Increase the cast range and duration of Holo Decoy by 100%.",
	},
	"NovaMasteryAmbushSnipe": {
		Name: "Ambush Snipe",
		Text: "Increases Snipe's range by 40% while Cloaked.",
	},
	"NovaMasteryCripplingShot": {
		Name: "Crippling Shot",
		Text: "Pinning Shot lowers a Hero's Armor by 25 for the duration of the slow, causing them to take 25% increased damage.",
	},
	"NovaMasteryDigitalShrapnel": {
		Name: "Digital Shrapnel",
		Text: "Digital Shrapnel \nReduces the cooldown of Holo Decoy by 2 seconds. Holo Decoy explodes on expiration, dealing 210 damage to nearby enemies.",
	},
	"NovaMasteryDoubleTap": {
		Name: "Double Tap",
		Text: "Pinning Shot now has 2 charges.",
	},
	"NovaMasteryExplosiveShot": {
		Name: "Explosive Round",
		Text: "Snipe also deals 50% damage to enemies near the impact.",
	},
	"NovaMasteryFastReload": {
		Name: "Fast Reload",
		Text: "Triple Tap's cooldown is reset if it kills an enemy Hero.",
	},
	"NovaMasteryLethalDecoy": {
		Name: "Lethal Decoy",
		Text: "Holo Decoy deals 40% of Nova's damage.",
	},
	"NovaMasteryLongshot": {
		Name: "Longshot",
		Text: "Increases the cast range of Pinning Shot by 30%. Pinning Shot also increases the range of Nova's next Basic Attack by 2.",
	},
	"NovaMasteryPerfectShotSnipe": {
		Name: "Perfect Shot",
		Text: "Hitting an enemy Hero with Snipe reduces the cooldown by 3 seconds. ",
	},
	"NovaMasteryPrecisionBarrage": {
		Name: "Precision Barrage",
		Text: "Precision Strike holds 2 charges with a 3 second cooldown between using each charge.",
	},
	"NovaMasteryPsiOpRangefinder": {
		Name: "Psi-Op Rangefinder",
		Text: "Increases Snipe's range by 20% and reduces the Cooldown by 2 seconds.",
	},
	"NovaMasteryPsionicEfficiency": {
		Name: "Psionic Efficiency",
		Text: "Removes Snipe's Mana cost and increases its range by 15%.",
	},
	"NovaMasteryTazerRounds": {
		Name: "Tazer Rounds",
		Text: "Increases the duration of Pinning Shot's slow to 4 seconds.",
	},
	"NovaNewRemoteDelivery": {
		Name: "Remote Delivery",
		Text: "Increases Holo Decoy cast range by 100% and its sight range by 50%",
	},
	"NovaRailgun": {
		Name: "Railgun",
		Text: "Snipe penetrates through the first enemy hit and deals 50% damage to subsequent targets. Snipe cooldown is reduced by 1 second for each target hit.",
	},
	"NovaRapidProjection": {
		Name: "Rapid Projection",
		Text: "Reduces Holo Decoy's cooldown and Mana cost by 40%.",
	},
	"NovaSnipeMaster": {
		Name: "Snipe Master",
		Text: "Hitting an enemy Hero with Snipe permanently increases the damage of Snipe by 15%. Stacks up to 5 times. All stacks are lost if Snipe fails to hit an enemy.",
	},
	"NovaTacticalEspionage": {
		Name: "Tactical Espionage",
		Text: "While Stealthed, Nova gains 4 Mana per second.",
	},
	"PlaceholderRangedCombatStyle": {
		Name: "Ranged Placeholder",
		Text: "Placeholder Combat Style\n+20% Basic Attack\n+1 Range",
	},
	"ProbiusAggressiveMatrixWarpInPylon": {
		Name: "Aggressive Matrix",
		Text: "Pylon's Power Field grants allied Heroes 35% increased Attack Damage.",
	},
	"ProbiusConstructAdditionalPylonsPylonOvercharge": {
		Name: "Construct Additional Pylons",
		Text: "Probius can now have up to 3 active Pylons, and increase the damage of Pylon Overcharge by 25%.",
	},
	"ProbiusEchoPulseDisruptionPulse": {
		Name: "Echo Pulse",
		Text: "Disruption Pulse now returns to Probius 1.2 seconds after reaching its target, dealing 75% damage on the return trip.",
	},
	"ProbiusGateKeeperNullGate": {
		Name: "Gate Keeper",
		Text: "Null Gate lasts indefinitely if either end is within a Power Field.\nOnly one Null Gate may be active at a time.",
	},
	"ProbiusGatherMineralsPhotonCannonQuest": {
		Name: "Gather Minerals",
		Text: "Quest: Enemy Heroes and Minions drop Minerals when killed. Collect them to increase the Health of Photon Cannons by 8, up to 560. \nReward: After collecting 70 Minerals, Photon Cannons deal 35% more damage, can see over obstacles, and reveal nearby Cloaked units.",
	},
	"ProbiusGravityWellWarpRift": {
		Name: "Gravity Well",
		Text: "Enemies are slowed more the closer they are to the center of the Warp Rift, up to a maximum slow of 60%.",
	},
	"ProbiusInterferenceWarpRift": {
		Name: "Interference",
		Text: "Enemy Heroes hit by Warp Rift detonations have their Spell Power reduced by 35% for 5 seconds.",
	},
	"ProbiusNullGateHeroic": {
		Name: "Null Gate",
		Text: "Vector Targeting\nProject a barrier of negative energy in the target direction that lasts 4 seconds. Enemies who touch the barrier take 68 damage per second and are slowed by 80% for as long as they remain in contact with it.",
	},
	"ProbiusParticleAcceleratorDisruptionPulse": {
		Name: "Particle Accelerator",
		Text: "Disruption Pulse deals 10% more damage for each enemy or Warp Rift hit, up to 50%.",
	},
	"ProbiusPhotonBarrierPhotonCannon": {
		Name: "Photon Barrier",
		Text: "While a Photon Cannon is alive and powered, Probius gains 30 Spell Armor.",
	},
	"ProbiusPowerOverflowingWarpInPylon": {
		Name: "Power Overflowing",
		Text: "Pylon's Power Field grants allied Heroes 10% increased Spell Power and 2 Mana per second.",
	},
	"ProbiusProbiusLoopWarpRift": {
		Name: "Probius Loop",
		Text: "Whenever a Rift explosion hits 2 or more enemy Heroes, create a new Warp Rift in the same location.",
	},
	"ProbiusPylonOverchargeHeroic": {
		Name: "Pylon Overcharge",
		Text: "Increase the size of Pylon power fields and allow them to attack enemies within it for 96 damage per second. Lasts 10 seconds.",
	},
	"ProbiusQuantumEntanglementWarpRift": {
		Name: "Quantum Entanglement",
		Text: "Enemies continue to be slowed for 3 seconds after their last contact with a Warp Rift.",
	},
	"ProbiusRepulsorWarpRift": {
		Name: "Repulsor",
		Text: "Enemies hit by a Warp Rift explosion are knocked away from the center.",
	},
	"ProbiusRiftShockWarpRift": {
		Name: "Rift Shock",
		Text: "Hitting an enemy Hero with Warp Rift explosion increases Probius's damage against them by 20% for 10 seconds.",
	},
	"ProbiusShieldBatteryWarpInPylon": {
		Name: "Shield Battery",
		Text: "Pylon's Power Field grants allied Heroes 28 Shields per second, up to 112. Shields persist for 2.2 seconds after exiting a Pylon Power Field.",
	},
	"ProbiusShieldCapacitor": {
		Name: "Shield Capacitor",
		Text: "Probius gains permanent Shields equal to 10% of his max Health. Shields regenerate quickly as long as he hasn't taken damage recently.",
	},
	"ProbiusShootEmUptoRiftsDisruptionPulse": {
		Name: "Shoot 'Em Up",
		Text: "Hitting a Warp Rift with Disruption Pulse causes 4 additional pulses that deal 50% damage to be fired from the impact location in different directions.  \nAdditional Pulses do not benefit from Echo Pulse or Particle Accelerator.",
	},
	"ProbiusTowerDefensePhotonCannon": {
		Name: "Tower Defense",
		Text: "When Photon Cannon damages an enemy Hero, its cooldown is reduced by 1.2 seconds.",
	},
	"ProbiusTurboChargedWorkerRush": {
		Name: "Turbo Charged",
		Text: "Worker Rush grants an additional 10% passive Movement Speed while in a Power Field, and its cooldown is reduced by 10 seconds.",
	},
	"ProbiusWarpResonanceWarpRiftQuest": {
		Name: "Warp Resonance",
		Text: "Quest: Hit Heroes with Warp Rift explosions.\nReward: After hitting 10 Heroes with Warp Rift explosions, increase the explosion damage by 100.\nReward: After hitting 20 Heroes with Warp Rift explosions, Warp Rift gains 1 additional charge.",
	},
	"RagnarosBlastWaveBlastEcho": {
		Name: "Blast Echo",
		Text: "After Blast Wave explodes, another Blast Wave is created on Ragnaros.",
	},
	"RagnarosBlastWaveEngulfingFlame": {
		Name: "Engulfing Flame",
		Text: "Increase Blast Wave's damage by 75% and its radius by 15%.",
	},
	"RagnarosBlastWaveSlowBurn": {
		Name: "Slow Burn",
		Text: "Blast Wave slows enemies hit by 40% for 2 seconds.",
	},
	"RagnarosBlastWaveSuperheated": {
		Name: "Superheated",
		Text: "When Blast Wave is used on Ragnaros, it deals 100% increased damage.",
	},
	"RagnarosBlastWaveTemperedFlame": {
		Name: "Tempered Flame",
		Text: "When Blast Wave damages an enemy Hero, gain a Shield equal to 100% of the damage dealt for 3 seconds.",
	},
	"RagnarosBlisteringAttacks": {
		Name: "Blistering Attacks",
		Text: "Every 10 seconds, Ragnaros's next Basic Attack against a Hero deals 60% bonus damage. Hitting enemies with Basic Abilities reduces this cooldown by 1 second.",
	},
	"RagnarosCatchingFire": {
		Name: "Catching Fire",
		Text: "Quest: Gathering a Regeneration Globe increases Ragnaros' Health Regeneration by 1.2 per second, up to 18.8.\nReward: After gathering 15 Regeneration Globes, activate Catching Fire to gain 25 Armor for 3 seconds, reducing damage taken by 25%.",
	},
	"RagnarosEmpowerSulfurasCauterizeWounds": {
		Name: "Cauterize Wounds",
		Text: "Empower Sulfuras heals for an additional 40% of damage dealt to Heroes over 2 seconds.",
	},
	"RagnarosEmpowerSulfurasGiantScorcher": {
		Name: "Giant Scorcher",
		Text: "Empower Sulfuras burns enemy Heroes for 9% of their maximum Health over 3 seconds. This additional damage does not heal Ragnaros.",
	},
	"RagnarosEmpowerSulfurasHandOfRagnaros": {
		Name: "Hand of Ragnaros",
		Text: "If Empower Sulfuras hits at least 2 enemy Heroes, it restores 10 Mana and its cooldown is reduced by 2 seconds.",
	},
	"RagnarosEmpowerSulfurasSulfurasHungers": {
		Name: "Sulfuras Hungers",
		Text: "Quest: Every time Empower Sulfuras kills a Minion, its damage is increased by 1, up to 30.\nReward: After killing 30 Minions, increase its damage by an additional 90.",
	},
	"RagnarosLavaWave": {
		Name: "Lava Wave",
		Text: "Release a wave of lava from Ragnaros's Core that travels down the targeted lane, dealing 240 damage per second to non-Structure enemies in its path and instantly killing enemy Minions. Damage increased by 100% versus Heroes.",
	},
	"RagnarosLavaWaveLavaSurge": {
		Name: "Lava Surge",
		Text: "Lava Wave gains an additional charge and its cooldown is reduced by 30 seconds.",
	},
	"RagnarosLivingMeteorFireWard": {
		Name: "Fire Ward",
		Text: "When Living Meteor hits an enemy Hero, gain a charge of Spell Block, reducing damage from the next enemy ability by 50%.\nStores up to 2 charges.",
	},
	"RagnarosLivingMeteorMeteorBomb": {
		Name: "Meteor Bomb",
		Text: "Living Meteor explodes at the end of its path, dealing 220 damage to all nearby enemies.",
	},
	"RagnarosLivingMeteorMoltenPower": {
		Name: "Molten Power",
		Text: "Each enemy Hero hit by Living Meteor increases the damage of the next Living Meteor by 20%, up to a maximum of 100%.",
	},
	"RagnarosLivingMeteorShiftingMeteor": {
		Name: "Shifting Meteor",
		Text: "Each time an enemy is hit by the same Living Meteor, they take 3% increased damage from it.\nQuest: Hit 75 Heroes with Living Meteor.\nReward: After hitting 75 Heroes with Living Meteor, its duration is increased to 2.2 seconds, and it can be reactivated to change its direction once per use.",
	},
	"RagnarosMoltenCoreHeroicDifficulty": {
		Name: "Heroic Difficulty",
		Text: "Molten Core has 25% increased Health and Damage, and its cooldown is reduced by 50 seconds.",
	},
	"RagnarosResilientFlame": {
		Name: "Resilient Flame",
		Text: "When Ragnaros is Stunned, he gains 40 Armor for 3 seconds, reducing damage taken by 40%. \nThis effect has a 15 second cooldown.",
	},
	"RagnarosSubmerge": {
		Name: "Submerge",
		Text: "Ragnaros submerges below, entering Stasis and healing for 600 Health over 3 seconds.",
	},
	"RagnarosSulfurasSmash": {
		Name: "Sulfuras Smash",
		Text: "Hurl Sulfuras at the target area, landing after 0.8 seconds, dealing 198 damage. Enemies in the center take 594 damage instead and are Stunned for 0.5 seconds.",
	},
	"RagnarosSulfurasSmashFlamesOfSulfuron": {
		Name: "Flames of Sulfuron",
		Text: "Sulfuras Smash slows enemies by 50% for 2 seconds, and enemies in the center are stunned for an additional 0.5 seconds.",
	},
	"RaynorACardToPlay": {
		Name: "A Card to Play",
		Text: "Whenever a Hero (ally or enemy) is killed, the cooldown of Raynor's Heroic Ability is reduced by 15 seconds.",
	},
	"RaynorHeroicAbilityHyperion": {
		Name: "Hyperion",
		Text: "Order the Hyperion to make a strafing run dealing 66 damage a second, hitting up to 4 enemies. Also occasionally fires its Yamato Cannon on Structures for 794 damage. Lasts 12 seconds.",
	},
	"RaynorHeroicAbilityRaynorsRaiders": {
		Name: "Raynor's Raiders",
		Text: "Summon two Stealthed Banshees that attack an enemy. Each Banshee deals 49.6 damage a second and lasts 22 seconds. Can reactivate the Ability to retarget the Banshees.",
	},
	"RaynorHyperionGroundStrafe": {
		Name: "Scorched Earth",
		Text: "An additional set of lasers blast the ground 5 times per second, dealing 40 damage in an area.",
	},
	"RaynorInspireConfidentAim": {
		Name: "Confident Aim",
		Text: "Lowers the cooldown of Penetrating Round by 4 seconds if it hits an enemy Hero.",
	},
	"RaynorInspireSteelResolve": {
		Name: "Steel Resolve",
		Text: "Increases Inspire's duration by 50% and causes Adrenaline Rush to also apply Inspire.",
	},
	"RaynorLeadFromTheFrontPuttinOnAClinic": {
		Name: "Puttin' On a Clinic",
		Text: "Whenever an enemy Minion Raynor has recently damaged is destroyed, his Ability cooldowns are reduced by 1.5 seconds.",
	},
	"RaynorMasteryActivatedRushAdrenalineRush": {
		Name: "Activated Rush",
		Text: "Lowers the cooldown of Adrenaline Rush by 10 seconds, and it can be manually activated.",
	},
	"RaynorMasteryBattleHyperionHyperion": {
		Name: "Battle Hyperion",
		Text: "Battle Hyperion \nNumber of shots on each target increased to 2 per volley. Every 4 seconds Hyperion will fire its Yamato cannon at enemy structures.",
	},
	"RaynorMasteryBullseyePenetratingRound": {
		Name: "Bullseye",
		Text: "The first enemy hit by Penetrating Round is stunned for 1 second.",
	},
	"RaynorMasteryClusterRoundPenetratingRound": {
		Name: "Cluster Round",
		Text: "Penetrating Round damage is increased by 20% for each additional target hit up to 100%, and the width is increased by 50%.",
	},
	"RaynorMasteryFightorFlightAdrenalineRush": {
		Name: "Fight or Flight",
		Text: "Whenever Adrenaline Rush activates, it gives 25 Armor, reducing damage taken by 25% for 4 seconds.  Adrenaline Rush can also be manually activated.",
	},
	"RaynorMasteryGiveMeMoreAdrenalineRush": {
		Name: "Give Me More!",
		Text: "Increases Adrenaline Rush healing by 50%.",
	},
	"RaynorMasteryHelsAngelsRaynorsBanshees": {
		Name: "Dusk Wings",
		Text: "Banshees remain Stealthed while attacking and fire 50% more frequently.",
	},
	"RaynorMasteryInspireRevolutionOverdrive": {
		Name: "Revolution Overdrive",
		Text: "Gain 10% Movement Speed while affected by Inspire.  Increase this bonus by 5% for each allied Hero nearby when Inspire is cast.",
	},
	"RaynorPenetratingRoundDoubleBarreled": {
		Name: "Double-Barreled",
		Text: "Penetrating Round gains a second charge.",
	},
	"RaynorPenetratingRoundHamstringShot": {
		Name: "Hamstring Shot",
		Text: "Enemies hit by Penetrating Round have a 20% Movement Speed slow for 3 seconds.",
	},
	"RaynorRelentlessLeader": {
		Name: "Relentless Leader",
		Text: "Being Stunned or Rooted knocks away nearby enemies.\nThis effect has a 8 second cooldown.\nPassive: The cooldown of Adrenaline Rush is reduced by 10 seconds.",
	},
	"RehgarEarthlivingEnchant": {
		Name: "Earthliving Enchant",
		Text: "When Chain Heal heals a Hero below 50% Health, they are healed an additional 200 Health over 5 seconds. ",
	},
	"RehgarGhostWolfBloodAndThunder": {
		Name: "Blood and Thunder",
		Text: "Ghost Wolf attacks reduce Basic Ability cooldowns by 2 seconds.",
	},
	"RehgarHeroicAbilityAncestralHealing": {
		Name: "Ancestral Healing",
		Text: "After 1 second, heal an allied Hero for 1552 Health.\nCannot be used on Rehgar.",
	},
	"RehgarHeroicAbilityBloodlust": {
		Name: "Bloodlust",
		Text: "Grant nearby allied Heroes 40% Attack Speed and 25% Movement Speed and causes them to heal for 30% of the Basic Attack damage to their primary target. Lasts for 8 seconds.",
	},
	"RehgarHungeroftheWolf": {
		Name: "Hunger of the Wolf",
		Text: "Ghost Wolf attacks against Heroes deal an additional 5% of the target's maximum Health and heal Rehgar for 5.1% of his maximum Health.",
	},
	"RehgarMasteryChainReaction": {
		Name: "Chain Reaction",
		Text: "Chain Heals on allies with Lightning Shield active are increased by 25%.",
	},
	"RehgarMasteryColossalTotem": {
		Name: "Colossal Totem",
		Text: "Increases the area and range of Earthbind Totem by 50%.",
	},
	"RehgarMasteryDeepHealing": {
		Name: "Deep Healing",
		Text: "Deep Healing \nIncreases Chain Heal's healing by 30% on targets that are below 50% health.",
	},
	"RehgarMasteryEarthGraspTotem": {
		Name: "Earthgrasp Totem",
		Text: "When Earthbind Totem is first cast, it slows nearby enemies by 90% for 1 second.",
	},
	"RehgarMasteryEarthShield": {
		Name: "Earth Shield",
		Text: "Lightning Shield gives a Shield that absorbs damage equal to 12% of their maximum Health for 3 seconds.",
	},
	"RehgarMasteryElectricCharge": {
		Name: "Electric Charge",
		Text: "Increases the radius of Lightning Shield by 25%.",
	},
	"RehgarMasteryFarseersBlessing": {
		Name: "Farseer's Blessing",
		Text: "Increases healing amount by 50%. Allies near the target are healed for 50% of the amount of health regained.",
	},
	"RehgarMasteryFarsight": {
		Name: "Farsight",
		Text: "Activate to reveal an area for 10 seconds.  Enemies in the area are revealed for 4 seconds.",
	},
	"RehgarMasteryFeralHeart": {
		Name: "Feral Heart",
		Text: "Increases Health and Mana Regeneration by 75% while in Ghost Wolf form.",
	},
	"RehgarMasteryGladiatorsWarShout": {
		Name: "Gladiator's War Shout",
		Text: "Increases Bloodlust's healing from 30% to 60% of Basic Attack damage done.",
	},
	"RehgarMasteryHealingSurge": {
		Name: "Healing Surge",
		Text: "Increases Chain Heal's healing on the primary target by 25% and heal an additional target.",
	},
	"RehgarMasteryLightningBond": {
		Name: "Lightning Bond",
		Text: "Casting Lightning Shield on an ally also casts an untalented version on Rehgar.",
	},
	"RehgarMasteryLightningTotem": {
		Name: "Lightning Totem",
		Text: "Lightning Totem \nEarthbind Totem automatically gains Lightning Shield when created.",
	},
	"RehgarMasteryShamanHealingWard": {
		Name: "Healing Totem",
		Text: "Activate to place a Totem that heals allies in an area for 1.9% of their maximum Health every second for 10 seconds.",
	},
	"RehgarMasterySpiritwalkersGrace": {
		Name: "Spiritwalker's Grace",
		Text: "Reduces Chain Heal's Mana cost from 65 to 45.",
	},
	"RehgarMasteryStormcaller": {
		Name: "Stormcaller",
		Text: "When Lightning Shield damages an enemy it restores 4 Mana, up to 40.",
	},
	"RehgarMasteryTidalWaves": {
		Name: "Tidal Waves",
		Text: "Reduces Chain Heal's cooldown by 1 second for each Hero healed.",
	},
	"RehgarMasteryWolfRun": {
		Name: "Wolf Run",
		Text: "Increases the Movement Speed of Ghost Wolf from 20% to 30%.",
	},
	"RehgarRisingStorm": {
		Name: "Rising Storm",
		Text: "Every time Lighting Shield damages an enemy Hero, increase that Lightning Shield's damage by 10%. Stacks up to 20 times.",
	},
	"RehgarTotemicProjection": {
		Name: "Totemic Projection",
		Text: "Increases Earthbind Totem's duration from 8 to 12 seconds, and it can be reactivated to move an existing totem to a new location once.",
	},
	"RexxarAnimalHusbandry": {
		Name: "Animal Husbandry",
		Text: "Rexxar and Misha gain 2 Maximum Health every second Rexxar remains alive. These bonuses are lost on Rexxar's death.",
	},
	"RexxarAspectOfTheHawkSpiritSwoop": {
		Name: "Aspect of the Hawk",
		Text: "When Spirit Swoop hits an enemy Hero, Rexxar gains 125% Attack Speed for 3 seconds. Misha's Basic Attacks increase the duration of this buff by 0.8 seconds.",
	},
	"RexxarAspectoftheBeastCharge": {
		Name: "Aspect of the Beast",
		Text: "Misha's Basic Attacks lower the cooldown of Misha, Charge! by 1.2 seconds.",
	},
	"RexxarBarkskin": {
		Name: "Barkskin",
		Text: "Misha gains 50 Spell Armor for 5 seconds after using Misha, Charge!, taking 50% less damage from Abilities.",
	},
	"RexxarBearNecessitiesCharge": {
		Name: "Bear Necessities",
		Text: "Misha's Basic Attacks against Minions heal Rexxar for 50% of the damage dealt. While Mend Pet is active, this bonus is increased to 100%.",
	},
	"RexxarCriticalCareMendPet": {
		Name: "Critical Care",
		Text: "Increases the healing of Mend Pet by 50% while Misha is under 50% health.",
	},
	"RexxarDireBeast": {
		Name: "Dire Beast",
		Text: "Rexxar and Misha's Basic Attacks increase the damage of the next Misha, Charge! by 40%, stacking up to a 400% bonus.",
	},
	"RexxarFeignDeath": {
		Name: "Feign Death",
		Text: "Rexxar fakes his death, becoming Invulnerable and untargetable for 5 seconds. During this time you control Misha.",
	},
	"RexxarFlare": {
		Name: "Flare",
		Text: "Fire a flare at an area, revealing it for 20 seconds.",
	},
	"RexxarFrenzyofKalimdor": {
		Name: "Frenzy of Kalimdor",
		Text: "Rexxar's Basic Attacks deal 10% more damage, and Misha's Basic Attacks slow the target by 20% for 1.2 seconds.",
	},
	"RexxarHardenedSkin": {
		Name: "Hardened Skin",
		Text: "Rexxar and Misha gain 75 Armor for 4 seconds, reducing damage taken by 75%.",
	},
	"RexxarHeroicAbilityBestialWrath": {
		Name: "Bestial Wrath",
		Text: "Increases Misha's Basic Attack damage by 150% for 12 seconds.",
	},
	"RexxarHeroicAbilityUnleashTheBoars": {
		Name: "Unleash the Boars",
		Text: "Release a herd of boars that track down all enemy Heroes in a direction, dealing 110 damage, revealing, and slowing enemies by 40% for 5 seconds.",
	},
	"RexxarHungryBear": {
		Name: "Hungry Bear",
		Text: "Misha's Basic Attacks heal her for 4.5% of her maximum Health.",
	},
	"RexxarHunterGatherer": {
		Name: "Hunter-Gatherer",
		Text: "Quest: Gathering a Regeneration Globe increases Rexxar's Health Regeneration by 1.2 per second, up to 25 per second. \nReward: After gathering 20 Globes, Rexxar and Misha gain 10 Armor.",
	},
	"RexxarMendPetBloodOfTheRhino": {
		Name: "Blood of the Rhino",
		Text: "Increases Mend Pet's duration by 5 seconds.",
	},
	"RexxarSpiritBond": {
		Name: "Spirit Bond",
		Text: "Increases the duration of Bestial Wrath by 50% and Misha's Basic Attacks heal Rexxar for 50% of her damage dealt during Bestial Wrath.",
	},
	"RexxarSpiritBondEasyPrey": {
		Name: "Easy Prey",
		Text: "Increases Misha's damage to Minions and Mercenaries by 150%, and Misha gains 50 Armor against Minions and Mercenaries, reducing damage taken by 50%.",
	},
	"RexxarSpiritBondGrizzledFortitude": {
		Name: "Grizzled Fortitude",
		Text: "Every 7 seconds, Rexxar and Misha gain 75 Physical Armor against the next enemy Hero Basic Attack, reducing the damage taken by 75%.\nStores up to 2 charges. Both Rexxar and Misha have their own charges.",
	},
	"RexxarSpiritBondPrimalIntimidation": {
		Name: "Primal Intimidation",
		Text: "Activate to slow the Attack Speed by 50% and Movement Speed by 20% of nearby Heroes and Summons near Rexxar and Misha for 2.5 seconds.\nPassive: Heroes and Summons that attack Rexxar or Misha have their Attack Speed slowed by 20% for 2.5 seconds.",
	},
	"RexxarSpiritBondWildfireBear": {
		Name: "Wildfire Bear",
		Text: "Misha deals 28 damage per second to nearby enemies.",
	},
	"RexxarSpiritSwoopBirdOfPrey": {
		Name: "Bird of Prey",
		Text: "Increases Spirit Swoop's damage by 150% to Minions.",
	},
	"RexxarSpiritSwoopCripplingTalons": {
		Name: "Crippling Talons",
		Text: "Increases Spirit Swoop's slow amount to 50% and its duration to 3 seconds.",
	},
	"RexxarSurvivalistTraining": {
		Name: "Survivalist Training",
		Text: "Regeneration Globes restore 100% more Mana.",
	},
	"RexxarTakingFlightSpiritSwoop": {
		Name: "Taking Flight",
		Text: "Increases Spirit Swoop's range by 20%. If Spirit Swoop hits an enemy Hero, the Mana cost is refunded and its cooldown is reduced by 2 seconds.",
	},
	"RexxarThrilloftheHunt": {
		Name: "Thrill of the Hunt",
		Text: "Rexxar's Basic Attacks increase both his and Misha's Movement Speed by 25% for 2 seconds.",
	},
	"RexxarUnleashTheBoarsKillCommand": {
		Name: "Kill Command",
		Text: "Unleash the Boars deals 50% more damage and roots for 1.5 seconds.",
	},
	"SamuroAdvancingStrikesDeflection": {
		Name: "Deflection",
		Text: "Advancing Strikes grants 25 Physical Armor against Hero Basic Attacks, reducing their damage by 25%.",
	},
	"SamuroBlademastersPursuit": {
		Name: "Blademaster’s Pursuit",
		Text: "Advancing Strikes grants an additional 15% bonus Movement Speed, and its duration is increased by 2 seconds.",
	},
	"SamuroBurningBlade": {
		Name: "Burning Blade",
		Text: "Critical Strikes unleash a burst of flame, dealing an additional 65% of Samuro's Basic Attack damage to the target and nearby enemies.",
	},
	"SamuroHarshWinds": {
		Name: "Harsh Winds",
		Text: "Attacking a Hero from Stealth causes them to take 30% increased damage from Samuro and his Images for 3 seconds.",
	},
	"SamuroHeroicAbilityBladestorm": {
		Name: "Bladestorm",
		Text: "Become an Unstoppable whirlwind of death, dealing 235 damage per second to nearby enemies for 4 seconds.",
	},
	"SamuroHeroicAbilityIllusionMaster": {
		Name: "Illusion Master",
		Text: "Switch places with the target Mirror Image, removing most negative effects from Samuro and the target Mirror Image.\nPassive: You can control Mirror Images separately or as a group, and they deal an additional 10% of Samuro's damage.",
	},
	"SamuroMercilessStrikes": {
		Name: "Merciless Strikes",
		Text: "Basic Attacks against slowed, rooted, or stunned targets are always Critical Strikes.",
	},
	"SamuroMirage": {
		Name: "Mirage",
		Text: "Mirror Image grants Samuro and his Images 40 Spell Armor against the next 2 sources of Spell damage, reducing their damage by 40%.",
	},
	"SamuroMirrorImageWayOfTheBlade": {
		Name: "Way of the Blade",
		Text: "Critical Strike now happens every 3rd Basic Attack and deals an additional 20% of Samuro's Basic Attack damage.",
	},
	"SamuroMirroredSteel": {
		Name: "Mirrored Steel",
		Text: "Basic Attacks against Heroes reduce the cooldown of Mirror Image by 1 second.",
	},
	"SamuroOverpoweringBlow": {
		Name: "Crushing Blows",
		Text: "Critical Strike now has two charges, and Critical Strikes deal an additional 25% of Samuro's Basic Attack damage.",
	},
	"SamuroPhantomPain": {
		Name: "Phantom Pain",
		Text: "Critical Strikes deal an additional 40% of Samuro's Basic Attack damage for each Image that is active.",
	},
	"SamuroPressTheAttack": {
		Name: "Press the Attack",
		Text: "While Advancing Strikes is active, Basic Attacks increase Samuro's Attack Speed by 15%, up to 60%.",
	},
	"SamuroShukuchi": {
		Name: "Shukuchi",
		Text: "Wind Walk teleports Samuro a short distance in the direction he's currently facing.",
	},
	"SamuroThreeBladeStyle": {
		Name: "Three Blade Style",
		Text: "Samuro's Images gain an additional 50% of his Health and last up to 36 seconds.",
	},
	"SamuroWayOfIllusion": {
		Name: "Way of Illusion",
		Text: " Quest: Every time one of Samuro's Images Critically Strikes a Hero, he gains 0.5 Attack Damage, up to 20.\nReward: After hitting 40 Heroes, he gains an additional 20 Attack Damage.",
	},
	"SamuroWayOfTheWind": {
		Name: "Way of the Wind",
		Text: " Quest: Attacking a Hero from Stealth increases the Movement Speed bonus of Wind Walk by 1.2%, up to 24.8%.\nReward: After attacking 20 Heroes, Wind Walk's Unrevealable duration is increased to 2.5 seconds.",
	},
	"SamuroWhirlwindStorm": {
		Name: "Dance Of Death",
		Text: "Mirror Images use Bladestorm when Samuro does.",
	},
	"SamuroWindStrider": {
		Name: "Wind Strider",
		Text: "Wind Walk's cooldown is reduced by 6 seconds.",
	},
	"SamuroWindwalkKawarimi": {
		Name: "Kawarimi",
		Text: "Wind Walk creates an Image at Samuro's location that will continue whatever he was doing.",
	},
	"SamuroWindwalkOneWithTheWind": {
		Name: "One With The Wind",
		Text: "While Stealthed, gain 60 Armor, taking 60% reduced damage from all sources.",
	},
	"SgtHammerAmbush": {
		Name: "Ambush",
		Text: "Stealth when entering Siege Mode. Sgt. Hammer's next Basic Attack from Siege Mode will deal 100% more damage. Lose Stealth when Basic Attacking, using an Ability, taking damage, or returning to Tank Mode.",
	},
	"SgtHammerBullheadMines": {
		Name: "Bullhead Mines",
		Text: "Middle Spider Mine knocks target back a short distance.",
	},
	"SgtHammerFirstStrike": {
		Name: "First Strike",
		Text: "Basic Attacks deal 25% more damage if Sgt. Hammer hasn't been attacked within the last 5 seconds.",
	},
	"SgtHammerGraduatingRange": {
		Name: "Graduating Range",
		Text: "While in Siege Mode, Sgt. Hammer's Attack range increases by 1.1 every 3 seconds, up to 5.5.",
	},
	"SgtHammerHeroicAbilityBluntForceGun": {
		Name: "Blunt Force Gun",
		Text: "Fire a missile across the battlefield, dealing 500 damage to enemies in its path.",
	},
	"SgtHammerHeroicAbilityNapalmStrike": {
		Name: "Napalm Strike",
		Text: "Deals 164 damage on impact, and leaves a napalm area that deals 50 damage per second. Lasts for 4 seconds.",
	},
	"SgtHammerHyperCoolingEngines": {
		Name: "Hyper-Cooling Engines",
		Text: "Reduce the cooldown of Thrusters by 15 seconds.",
	},
	"SgtHammerMasteryAdvancedArtillery": {
		Name: "Advanced Artillery",
		Text: "Increase the damage bonus to long distance enemies by 10%.",
	},
	"SgtHammerMasteryAdvancedLavaStrikeNapalmStrike": {
		Name: "Advanced Lava Strike",
		Text: "Napalm Strike's range is increased by 75% and its impact does 50% more damage.",
	},
	"SgtHammerMasteryExcessiveForceConcussiveBlast": {
		Name: "Excessive Force",
		Text: "Double the knock back distance.",
	},
	"SgtHammerMasteryFlakCannons": {
		Name: "Barricade",
		Text: "Create a wall of path blocking debris for 4 seconds.",
	},
	"SgtHammerMasteryHoverSiegeMode": {
		Name: "Hover Siege Mode",
		Text: "Sgt. Hammer can move at 50% Movement Speed in Siege Mode.",
	},
	"SgtHammerMasteryLethalBlastConcussiveBlast": {
		Name: "Lethal Blast",
		Text: "Increase the damage of Concussive Blast by 50%.",
	},
	"SgtHammerMasteryMaelstromShells": {
		Name: "Maelstrom Shells",
		Text: "Increase standard Basic Attack range by 20%.",
	},
	"SgtHammerMasteryMineField": {
		Name: "Mine Field",
		Text: "Increase the number of mines by 2.",
	},
	"SgtHammerMasteryOrbitalBFGBluntForceGun": {
		Name: "Orbital BFG",
		Text: "Blunt Force Gun's missile orbits the planet every 5 seconds. Only the last missile fired orbits.",
	},
	"SgtHammerMasterySlowingMinesSpiderMines": {
		Name: "Slowing Mines",
		Text: "Increase the Movement Speed slow of Spider Mines to 40%, and the duration to 2 seconds.",
	},
	"SgtHammerResistant": {
		Name: "Resistant",
		Text: "Being Stunned or Rooted while in Siege Mode grants 50 Armor for 3 seconds, reducing damage taken by 50%.",
	},
	"SonyaTalentIgnorePain": {
		Name: "Ignore Pain",
		Text: "Activate to gain 75 Armor for 4 seconds, taking 75% less damage.\nUsable while Whirlwinding.",
	},
	"SonyaTalentNervesOfSteel": {
		Name: "Nerves of Steel",
		Text: "Activate to gain 30% of your maximum Health as a Shield for 5 seconds. Usable while Whirlwinding.",
	},
	"StitchesCannibalize": {
		Name: "Cannibalize",
		Text: "Basic Attacks against Heroes heal Stitches for 5.1% of his maximum Health.",
	},
	"StitchesCombatStyleTenderizer": {
		Name: "Tenderizer",
		Text: "Basic Attacks slow enemy Movement Speed by 25% for 1.5 seconds.",
	},
	"StitchesCombatStyleVileCleaver": {
		Name: "Vile Cleaver",
		Text: "Basic Attacks create a cloud of Vile Gas on the target.",
	},
	"StitchesFleaBag": {
		Name: "Flea Bag",
		Text: "Whenever Stitches is stunned or rooted, his Basic Ability cooldowns are reduced by 6 seconds.",
	},
	"StitchesHeroicAbilityGorge": {
		Name: "Gorge",
		Text: "Consume an enemy Hero, trapping them for 4 seconds. When Gorge ends, the enemy Hero takes 274 damage. The trapped Hero cannot move or act and doesn't take damage from other sources. \nCannot be used on massive Heroes.",
	},
	"StitchesHeroicAbilityPutridBile": {
		Name: "Putrid Bile",
		Text: "Emit bile that deals 37 damage per second to enemies within, slowing them by 35%. Stitches gains 20% Movement Speed while emitting bile. Lasts 8 seconds.",
	},
	"StitchesHungryforMore": {
		Name: "Hungry for More",
		Text: "Quest: Gathering a Regeneration Globe increases Stitches's maximum Health by 30, up to 900. \nReward: Upon gathering 30 Globes, increase Movement Speed by 10%.",
	},
	"StitchesMasteryChewYourFood": {
		Name: "Chew Your Food",
		Text: "Using Devour also heals Stitches for 9.6% of his max Health over 3.1 seconds.",
	},
	"StitchesMasteryFishingHook": {
		Name: "Fishing Hook",
		Text: "Hook has an additional 40% range.",
	},
	"StitchesMasteryHeavySlam": {
		Name: "Heavy Slam",
		Text: "Slam damage increased by 40%.",
	},
	"StitchesMasteryHungryHungryStitchesGorge": {
		Name: "Hungry Hungry Stitches",
		Text: "For 4 seconds after using Gorge, Stitches can teleport a short distance to additional targets and Gorge them.",
	},
	"StitchesMasteryIndigestionDevour": {
		Name: "Indigestion",
		Text: "Using Devour also creates a Retchling that applies Vile Gas Poison when it attacks.",
	},
	"StitchesMasteryLastBiteDevour": {
		Name: "Last Bite",
		Text: "If an enemy dies within 3 seconds of being damaged by Devour, its cooldown is reduced by 12 seconds.",
	},
	"StitchesMasteryMegaSmashSlam": {
		Name: "Mega Smash",
		Text: "Range and arc of Slam increased by 25%.",
	},
	"StitchesMasteryPulverizeSlam": {
		Name: "Pulverize",
		Text: "Decreases Slam's cooldown by 2 seconds and it also slows enemies by 75% for 1 second.",
	},
	"StitchesMasteryPutridGroundSlam": {
		Name: "Putrid Ground",
		Text: "Enemies hit by Slam are infected with Vile Gas.",
	},
	"StitchesMasterySavorTheFlavorDevour": {
		Name: "Savor the Flavor",
		Text: "Quest: Using Devour on an enemy Hero permanently increases Stitches's Health Regeneration by 2 per second.",
	},
	"StitchesMasteryShishKabobHook": {
		Name: "Shish Kabob",
		Text: "Hook can pull up to 2 targets.",
	},
	"StitchesPotentBile": {
		Name: "Potent Bile",
		Text: "Putrid Bile lasts 6 seconds longer and its slow amount is increased from 35% to 45%.",
	},
	"StitchesRestorativeFumes": {
		Name: "Restorative Fumes",
		Text: "Whenever Vile Gas damages an enemy Hero, Stitches is healed for 27 Health.",
	},
	"StitchesToxicGas": {
		Name: "Toxic Gas",
		Text: "Increases Vile Gas' radius by 25% and its damage by 50%.",
	},
	"StukovBioExplosionSwitch": {
		Name: "Bio-Explosion Switch",
		Text: "Bio-Kill Switch now also detonates Lurking Arm, applying a Weighted Pustule to enemy Heroes inside and Silencing them for 2 seconds.",
	},
	"StukovBioticArmor": {
		Name: "Biotic Armor",
		Text: "Healing Pathogen grants 50 Physical Armor to its initial target while active on them.",
	},
	"StukovControlledChaos": {
		Name: "Controlled Chaos",
		Text: "Flailing Swipe gains 2 additional charges, but each use only swings 1 time, at maximum range.  Additionally, its cooldown is decreased to 30 seconds, and its Mana cost is reduced from 100 to 50.",
	},
	"StukovEyeInfection": {
		Name: "Eye Infection",
		Text: "Increase the final damage of Weighted Pustule by 50%. Additionally, detonating a Weighted Pustule with Bio-Kill Switch Blinds the target for 3 seconds.",
	},
	"StukovFetidTouch": {
		Name: "Fetid Touch",
		Text: "Quest: Hit Heroes with Weighted Pustule.\nReward: After hitting 15 Heroes, reduce the cooldown of Weighted Pustule by 2.5 seconds.\nReward: After hitting 30 Heroes, reduce the cooldown of Weighted Pustule by an additional 2.5 seconds and remove its Mana cost.",
	},
	"StukovGrowingInfestation": {
		Name: "Growing Infestation",
		Text: "Lurking Arm's area expands by 50% over 2.5 seconds but its cooldown is increased by 4 seconds.",
	},
	"StukovHeroicAbilityFlailingSwipe": {
		Name: "Flailing Swipe",
		Text: "Swipe 3 times in front of Stukov over 1.8 seconds, dealing 50 damage to enemies hit and knocking them away. Each swipe is larger than the previous.",
	},
	"StukovHeroicAbilityMassiveShove": {
		Name: "Massive Shove",
		Text: "Extend Stukov's arm. If it hits an enemy Hero, they are rapidly shoved until they collide with terrain, dealing 200 damage and Stunning them for 0.5 seconds. Stukov gains 50 Armor while shoving an enemy.",
	},
	"StukovItHungers": {
		Name: "It Hungers",
		Text: "Each time an enemy Hero is hit by Lurking Arm, reduce its cooldown by 0.5 seconds, up to 6.5 seconds, and restore 5 Mana.",
	},
	"StukovLingeringSpines": {
		Name: "Lingering Spines",
		Text: "Lurking Arm persists for 1.5 seconds after it is canceled.",
	},
	"StukovLongPitch": {
		Name: "The Long Pitch",
		Text: "Increase the range of Weighted Pustule by 100%.",
	},
	"StukovLowBlow": {
		Name: "Low Blow",
		Text: "Lurking Arm deals 150% more damage to enemy Heroes below 30% Health.",
	},
	"StukovOneGoodSpread": {
		Name: "One Good Spread...",
		Text: "After a Healing Pathogen infests 3 targets, restore 40 Mana and reduce the cooldown of Healing Pathogen by 2 seconds.",
	},
	"StukovPoppinPustules": {
		Name: "Poppin' Pustules",
		Text: "Increase the final damage of Weighted Pustule by 50%.\nQuest: Detonate 10 Weighted Pustules without dying.\nReward: Detonating a Weighted Pustule now applies its damage and Slow in an area.",
	},
	"StukovPoxPopuli": {
		Name: "Pox Populi",
		Text: "Bio-Kill Switch no longer removes Healing Pathogen but instead sets its duration to 3 seconds.",
	},
	"StukovPushComesToShove": {
		Name: "Push Comes To Shove",
		Text: "Massive Shove travels 25% faster. If Massive Shove pushes a target for more than 1.2 seconds, its cooldown is reduced by 15 seconds.",
	},
	"StukovSpineLauncher": {
		Name: "Spine Launcher",
		Text: "Basic Attacks become ranged and Slow enemies by 20% for 1.5 seconds, but deal 40% less damage.",
	},
	"StukovSuperstrain": {
		Name: "Superstrain",
		Text: "Whenever an ally with Healing Pathogen is Stunned or Rooted, they are instantly healed for 250 Health.",
	},
	"StukovTargetedExcision": {
		Name: "Targeted Excision",
		Text: "Detonating exactly 1 Weighted Pustule (but any number of Healing Pathogens) with Bio-Kill Switch reduces the cooldown of Bio-Kill Switch to 5 seconds.",
	},
	"StukovTopOff": {
		Name: "Top Off",
		Text: "Healing Pathogen's heal over time on its initial target is increased by 50% while they are above 60% Health.",
	},
	"StukovUniversalCarrier": {
		Name: "Universal Carrier",
		Text: "Healing Pathogen can continually spread through Stukov, but its healing is reduced by 50%.",
	},
	"StukovVigorousReuptake": {
		Name: "Vigorous Reuptake",
		Text: "Quest: Detonate 50 Healing Pathogens with Bio-Kill Switch.\nReward: Increase Bio-Kill Switch's healing by 40%.",
	},
	"StukovVirulentReaction": {
		Name: "Virulent Reaction",
		Text: "Detonating a Weighted Pustule on an enemy who is inside of Lurking Arm Roots them for 2 seconds.",
	},
	"StukovWeightedPustuleReactiveBallistospores": {
		Name: "Reactive Ballistospores",
		Text: "When taking damage below 50% Health, instantly spread a Weighted Pustule to all nearby enemy Heroes and reset the cooldown of Bio-Kill Switch.\nThis effect has a 10 second cooldown.",
	},
	"StukovWithinMyReach": {
		Name: "Within My Reach",
		Text: "Increase the range of Lurking Arm by 50%.",
	},
	"SylvanasDreadfulWake": {
		Name: "Dreadful Wake",
		Text: "Haunting Wave's application of Black Arrow on Minions and Mercenaries increases the duration to 7 seconds. Haunting Wave restores 5 Mana per Minion or Mercenary hit.",
	},
	"SylvanasHeroicAbilityMindControl": {
		Name: "Mind Control",
		Text: "After a 1 second cast, take control of an enemy Hero's movement, Silencing them and prevent them from doing anything else. Channel lasts 2.5 seconds.",
	},
	"SylvanasHeroicAbilityPossession": {
		Name: "Possession",
		Text: "Force an enemy Minion to fight for Sylvanas's team.\nStores up to 3 charges.",
	},
	"SylvanasHeroicAbilityWailingArrow": {
		Name: "Wailing Arrow",
		Text: "Shoot an arrow that can be reactivated to deal 228 damage and silence enemies in an area making them unable to use Abilities for 2.5 seconds. The arrow detonates automatically if it reaches maximum range.",
	},
	"SylvanasTalentBlackArrowsParalysis": {
		Name: "Paralysis",
		Text: "Increases Black Arrows' duration by 0.8 seconds.",
	},
	"SylvanasTalentColdEmbrace": {
		Name: "Cold Embrace",
		Text: "Shadow Dagger lowers the Armor of the initial enemy by 25 for 2 seconds, causing them to take 25% more damage.",
	},
	"SylvanasTalentCorruption": {
		Name: "Corruption",
		Text: "Basic Attacks against Structures destroy 2 Ammunition.",
	},
	"SylvanasTalentDarkLadysCallMindControl": {
		Name: "Dark Lady's Call",
		Text: "Mind Controlled enemies gain 100% Movement Speed for the duration.",
	},
	"SylvanasTalentLifeDrain": {
		Name: "Life Drain",
		Text: "Shadow Dagger heals Sylvanas for 37 each time it spreads to an enemy.",
	},
	"SylvanasTalentLostSoul": {
		Name: "Lost Soul",
		Text: "Reduce Shadow Dagger's cooldown by 1.2 seconds each time it spreads to a Hero.",
	},
	"SylvanasTalentMercenaryQueen": {
		Name: "Mercenary Queen",
		Text: "Friendly non-Boss Mercenaries near Sylvanas deal 60% more damage.",
	},
	"SylvanasTalentOverflowingQuiver": {
		Name: "Overflowing Quiver",
		Text: "When a nearby enemy Minion dies, a free Withering Fire is shot. This cannot hit Heroes.",
	},
	"SylvanasTalentOverwhelmingAffliction": {
		Name: "Overwhelming Affliction",
		Text: "Black Arrows now also applies to Heroes, slowing their Movement Speed by 6% for the duration. Stacks up to 5 times.",
	},
	"SylvanasTalentRemorseless": {
		Name: "Remorseless",
		Text: "After using an Ability, Sylvanas's next Basic Attack within 3 seconds deals 40% additional damage.",
	},
	"SylvanasTalentUnstablePoison": {
		Name: "Unstable Poison",
		Text: "Minions that die under the effects of Black Arrows explode, dealing 116 damage to nearby enemies.  Does not damage Heroes or Structures.",
	},
	"SylvanasTalentWillOfTheForsaken": {
		Name: "Will of the Forsaken",
		Text: "Activate to become Unstoppable and gain 40% Movement Speed for 3 seconds.",
	},
	"SylvanasTalentWindrunnerHauntingWave": {
		Name: "Windrunner",
		Text: "Teleporting with Haunting Wave fully recharges Withering Fire. Haunting Wave can be cast a second time for free within 5 seconds after teleporting.",
	},
	"SylvanasTalentWithTheWind": {
		Name: "With the Wind",
		Text: "Increases Withering Fire's range by 25%.",
	},
	"SylvanasTalentWitheringBarrage": {
		Name: "Withering Barrage",
		Text: "Increase Withering Fire charges by 1. Charges can be fired 33% faster.",
	},
	"SylvanasTalentWitheringFireBarbedShot": {
		Name: "Barbed Shot",
		Text: "Withering Fire deals 200% bonus damage to Minions, Mercenaries, and Monsters.",
	},
	"SylvanasTalentWitheringFireEvasiveFire": {
		Name: "Evasive Fire",
		Text: "When an enemy is hit with Withering Fire, gain 10% Movement Speed for 2 seconds, stacking up to 30%.",
	},
	"SylvanasWailingArrowDeafeningBlast": {
		Name: "Deafening Blast",
		Text: "Enemies at the center of Wailing Arrow's explosion take 50% more damage and are silenced for twice as long.",
	},
	"TalentNullificationShield": {
		Name: "Nullification Shield",
		Text: "Activate to gain 60 Spell Armor for 5 seconds.",
	},
	"TassadarAdunsWisdom": {
		Name: "Adun's Wisdom",
		Text: "While Oracle is active, Basic Abilities cooldown 50% faster.",
	},
	"TassadarDeepShift": {
		Name: "Deep Shift",
		Text: "Dimensional Shift's duration and Movement Speed bonuses are increased by 50%.",
	},
	"TassadarFocusedBeam": {
		Name: "Focused Beam",
		Text: "While Oracle is active, Basic Attacks against Heroes deal an additional 1% of the target's maximum Health as damage.",
	},
	"TassadarHeroicAbilityArchon": {
		Name: "Phase Shift",
		Text: "Tassadar transforms into an Archon and gains a Plasma Shield. His Basic Attacks deal 166 damage, slow the target by 30% for 1 second and splash for 83 damage to enemies within 2.5 range. Lasts for 10.5 seconds.\nPassive: Archon refreshes the cooldown of Dimensional Shift.",
	},
	"TassadarHeroicAbilityForceWall": {
		Name: "Force Wall",
		Text: "Create a wall that blocks all units from moving through it for 2 seconds. \nPassive: Increases the slow amount of Distortion Beam to 30%.",
	},
	"TassadarKhalasCelerityPlasmaShield": {
		Name: "Khala's Celerity",
		Text: "Plasma Shield also grants 20% Movement Speed for 4 seconds.",
	},
	"TassadarKhalasEmbrace": {
		Name: "Khala's Embrace",
		Text: "Increase the Life Steal of Plasma Shield to 75%.",
	},
	"TassadarKhalasLight": {
		Name: "Khala's Light",
		Text: "Upon expiration or breaking, the target of Plasma Shield gains 15 Armor for 3 seconds, causing them to take 15% less damage.",
	},
	"TassadarKhaydarinResonance": {
		Name: "Khaydarin Resonance",
		Text: "Quest: Regeneration Globes grant an additional 100% Mana. \nReward: After collecting 15 Regeneration Globes, Plasma Shield's value is increased by 20%.\nReward: After collecting 30 Regeneration Globes, Plasma Shield's value is increased by 40%.",
	},
	"TassadarMasteryForceBarrier": {
		Name: "Force Barrier",
		Text: "Force Wall's Range is increased by 100% and its cooldown is reduced by 5 seconds.",
	},
	"TassadarMasteryMentalAcuity": {
		Name: "Mental Acuity",
		Text: "Quest: Each Takedown earned lowers the cooldown of Oracle by 3 seconds, to a maximum of 15 seconds.\nReward: Upon getting 5 Takedowns, the sight range of Oracle is increased by 50%.",
	},
	"TassadarMasteryTwilightArchon": {
		Name: "Phase Shift",
		Text: "Basic Attacks while in Archon form increase its duration by 2 seconds.",
	},
	"TassadarNullification": {
		Name: "Nullification",
		Text: "Activate to reduce the damage the target enemy Hero deals by 75% for 4 seconds.",
	},
	"TassadarPhaseDisruption": {
		Name: "Phase Disruption",
		Text: "Enemy Heroes being hit by Psionic Storm have their Physical Armor reduced by 25, causing them to take 25% more damage from Basic Attacks.",
	},
	"TassadarPrismaticLink": {
		Name: "Prismatic Link",
		Text: "Basic Attacks bounce to hit 2 additional targets for 40% damage.",
	},
	"TassadarPsionicEcho": {
		Name: "Psionic Echo",
		Text: "A second Psionic Storm can be cast within 2 seconds of casting the first. Enemies can only be affected by one Psionic Storm at a time.",
	},
	"TassadarPsionicProjection": {
		Name: "Psionic Projection",
		Text: "Increases the cast range of Plasma Shield and Psionic Storm by 40%.",
	},
	"TassadarPsionicStormPsiInfusion": {
		Name: "Psi-Infusion",
		Text: "Quest: Gain 1 Mana for every non-Structure enemy hit by Psionic Storm.\nReward: After hitting 500 enemies, increase the size of Psionic Storm by 20%.\nReward: After hitting 1000 enemies, increase the damage of Psionic Storm by 20%",
	},
	"TassadarResonation": {
		Name: "Resonation",
		Text: "Distortion Beam slows enemies affected by Psionic Storm by 50%.",
	},
	"TassadarShieldBattery": {
		Name: "Shield Battery",
		Text: "Activate to give all nearby allied Heroes an untalented Plasma Shield.",
	},
	"TassadarTemplarsWill": {
		Name: "Templar's Will",
		Text: "Quest: Attacking enemy Heroes restores 8 mana per second.\nReward: After spending 30 seconds attacking Heroes, Distortion Beam's damage is increased by 125%.\nReward: After spending 60 seconds attacking Heroes, Distortion Beam's range is increased by an additional 1.",
	},
	"ThrallAlphaWolf": {
		Name: "Alpha Wolf",
		Text: "For 3 seconds after a Hero is hit by Feral Spirit, Thrall's Basic Attacks against them deal an additional 3% of the target's maximum Health as damage.\nPassive: Increase the duration of Feral Spirit's Root to 1.5 seconds.",
	},
	"ThrallAncestralWrath": {
		Name: "Ancestral Wrath",
		Text: "Activate to consume 8 stacks of Ancestral Wrath, damaging a target enemy Hero for 15% of their maximum Health over 3 seconds, and healing Thrall for 150% of the damage dealt.\nGain 1 stack of Ancestral Wrath every time Frostwolf Resilience activates.",
	},
	"ThrallCrashLightning": {
		Name: "Crash Lightning",
		Text: "Repeatable Quest: Hitting at least 2 Heroes with a single use of Chain Lightning increases the damage of its bounces by 12, up to 480.\nReward: After hitting multiple Heroes 20 times, Chain Lightning prioritizes bouncing to Heroes.",
	},
	"ThrallEchooftheElements": {
		Name: "Echo of the Elements",
		Text: "Quest: Kill Minions or Heroes within 1.5 seconds of hitting them with Chain Lightning.\nReward: After killing 10 enemies, reduce the Mana cost of Chain Lightning from 40 to 25.\nReward: After killing 20 enemies, gain a second charge of Chain Lightning.",
	},
	"ThrallFeralResilience": {
		Name: "Feral Resilience",
		Text: "Heroes damaged by Feral Spirit grant 3 stacks of Frostwolf Resilience and 2 stacks of Feral Resilience, granting 50 Physical Armor against the next incoming Hero Basic Attack.\nStacks are consumed on the next cast of Feral Spirit.",
	},
	"ThrallFrostwolfPack": {
		Name: "Frostwolf Pack",
		Text: "Quest: Hit 7 Heroes with Feral Spirit. Progress is lost on death and when Thrall fails to hit a Hero.\nReward: Reduce the cooldown and Mana cost of Feral Spirit by 50%.",
	},
	"ThrallHeroicAbilityEarthquake": {
		Name: "Earthquake",
		Text: "After 0.5 seconds, summon a massive Earthquake that pulses every 4 seconds. Each pulse lasts 2 seconds, Slowing all enemies in the area by 50%, and deals 50 damage to enemy Heroes. Does 3 pulses.",
	},
	"ThrallHeroicAbilitySundering": {
		Name: "Sundering",
		Text: "After 0.5 seconds, sunder the earth in a long line, dealing 290 damage and shoving enemies to the side, Stunning them for 1 second.",
	},
	"ThrallMaelstromWeapon": {
		Name: "Maelstrom Weapon",
		Text: "Repeatable Quest: Basic Attacks against Heroes while Windfury's Movement Speed bonus is active increase Attack Damage by 1.\nReward: After gaining 20 Attack Damage, increase the Movement Speed bonus of Windfury to 40%.\nReward: After gaining 40 Attack Damage, Thrall permanently gains 15% increased Movement Speed.",
	},
	"ThrallMasteryEarthenShields": {
		Name: "Earthen Shields",
		Text: "Allies within the Earthquake area gain a Shield equal to 15% of their maximum Health each pulse. This shield lasts 4 seconds.",
	},
	"ThrallMasteryFrostwolfsGrace": {
		Name: "Frostwolf's Grace",
		Text: "Frostwolf Resilience can be activated to instantly heal for 150% of its normal amount.",
	},
	"ThrallMasteryGraceOfAir": {
		Name: "Grace Of Air",
		Text: "Windfury attacks grant twice as many stacks of Frostwolf Resilience.",
	},
	"ThrallMasteryManaTide": {
		Name: "Mana Tide",
		Text: "Frostwolf Resilience restores 15 Mana and reduces Basic Ability cooldowns by 0.5 seconds.",
	},
	"ThrallMasteryRollingThunder": {
		Name: "Rolling Thunder",
		Text: "For 10 seconds after hitting an enemy with Chain Lightning, Thrall's next Basic Attack against them restores 15 Mana.\nPassive: Increase the range of Chain Lightning by 30% and the number of bounces by 1.",
	},
	"ThrallMasteryStoneWolves": {
		Name: "Stone Wolves",
		Text: "Increases the duration of Feral Spirit's root from 1 second to 1.5 seconds.",
	},
	"ThrallMasteryTempestFury": {
		Name: "Tempest Fury",
		Text: "The final strike of Windfury hits 3 times for 75% normal damage.",
	},
	"ThrallMasteryWorldbreaker": {
		Name: "Worldbreaker",
		Text: "Lower the cooldown of Sundering by 20 seconds. After Sundering impacts, it leaves behind an impassable rift, blocking unit movement for 2.6 seconds.",
	},
	"ThrallSpiritShield": {
		Name: "Spirit Shield",
		Text: "Every 45 seconds, gain 50 Spell Armor against the next enemy Ability and subsequent Abilities for 1.5 seconds. Frostwolf Resilience reduces this cooldown by 8 seconds.\nCan be toggled to allow or prevent this talent from triggering automatically.",
	},
	"ThrallThunderstorm": {
		Name: "Thunderstorm",
		Text: "Casting Chain Lightning on a Hero other than the last Hero it was cast on causes it to Slow targets by 8% for 2 seconds.\nRepeatable Quest: Each consecutive cast against a new Hero increases this bonus by 8%, up to 40%. Bonuses are reset if Thrall dies or casts Chain Lightning on the same Hero twice.\nReward: While at a 40% bonus, Chain Lightning's damage is increased by 25%.",
	},
	"ThrallWindRush": {
		Name: "Wind Rush",
		Text: "Activate to immediately teleport to a target location and gain Windfury.",
	},
	"ThrallWindStalker": {
		Name: "Wind Stalker",
		Text: "Quest: Attack enemies with Windfury active. \nReward: After attacking 80 enemies, the movement speed bonus of Windfury is increased from 30% to 40%.",
	},
	"TinkerCombatStyleBreakitDown": {
		Name: "Break it Down!",
		Text: "Scrap causes Abilities to cooldown four times as fast for 3.1 seconds.",
	},
	"TinkerCombatStyleClockwerkSteamFists": {
		Name: "Clockwerk Steam Fists",
		Text: "Quest: Basic Attacks increase the duration of active Rock-It! Turrets by 1.5 seconds.\nReward: After increasing turret duration by 600 total seconds, reduce the charge cooldown of Rock-It! Turret by 3 seconds.",
	},
	"TinkerCombatStyleReduceReuseRecycle": {
		Name: "Reduce, Reuse, Recycle",
		Text: "Enemy Minions that die nearby have a 15% chance to drop Scrap.",
	},
	"TinkerCombatStyleScrapoMaticSmelter": {
		Name: "Scrap-o-Matic Smelter",
		Text: "Increases the Mana restored from Scrap by 100%.",
	},
	"TinkerGoblinRepairs": {
		Name: "Goblin Repairs",
		Text: "Quest: Gathering a Regeneration Globe increases Health Regeneration by 1 per second, up to 25.\nReward: After gathering 25 Globes, gain an additional 15 Health Regeneration per second.",
	},
	"TinkerHeroicAbilityGravOBomb3000": {
		Name: "Grav-O-Bomb 3000",
		Text: "After a 2 second delay, pull enemies toward the center of an area and deal 251 damage.",
	},
	"TinkerHeroicAbilityRoboGoblin": {
		Name: "Robo-Goblin",
		Text: "Activate to gain 30 Armor and 30% Movement Speed for 4 seconds. \nPassive: Basic Attacks deal 100% bonus damage. ",
	},
	"TinkerHiredGoons": {
		Name: "Hired Goons",
		Text: "Friendly non-Boss Mercenaries near Gazlowe deal 50% more damage. Rock-it! Turrets gain 50 Armor against Minions and Mercenaries, reducing their damage taken by 50%.",
	},
	"TinkerItsRainingScrap": {
		Name: "It's Raining Scrap",
		Text: "Activate to create 5 Scrap around Gazlowe over 5 seconds.",
	},
	"TinkerMasteryARKReaktor": {
		Name: "ARK Reaktor",
		Text: "Reduces Xplodium Charge's cooldown by 2 seconds.  If Xplodium Charge hits an enemy Hero, gain 2 charges of Rock-It! Turret.",
	},
	"TinkerMasteryEZPZDimensionalRipper": {
		Name: "EZ-PZ Dimensional Ripper",
		Text: "Deth Lazor Slows Heroes by 50% and Stuns non-Heroes for 3 seconds.",
	},
	"TinkerMasteryEngineGunk": {
		Name: "Engine Gunk",
		Text: "Rock-It! Turret attacks slow enemies by 20% for 2 seconds.",
	},
	"TinkerMasteryExtraTNT": {
		Name: "Extra TNT",
		Text: "Increases Xplodium Charge damage by 10% per target hit, up to 100%.",
	},
	"TinkerMasteryGoblinFusion": {
		Name: "Goblin Fusion",
		Text: "Deth Lazor can be charged for 1 additional second to increase its damage bonus by an additional 100%.",
	},
	"TinkerMasteryHyperfocusCoils": {
		Name: "Hyperfocus Coils",
		Text: "The mana cost of Deth Lazor is reduced by 30 and it charges twice as fast.",
	},
	"TinkerMasteryLongRangedTurrets": {
		Name: "Long-Ranged Turrets",
		Text: "Turret range increased by 40%.",
	},
	"TinkerMasteryMechaLord": {
		Name: "Mecha-Lord",
		Text: "Basic Attacks deal an additional 100% damage.\nPassive: Gain 25 Armor.",
	},
	"TinkerMasteryMiniatureBlackHole": {
		Name: "Miniature Black Hole",
		Text: "Increases Grav-O-Bomb radius by 25% and damage by 50%.",
	},
	"TinkerMasteryQuikReleaseCharge": {
		Name: "Kwik Release Charge",
		Text: "Xplodium Charge gains a second charge.",
	},
	"TinkerMasteryRockitTurretXL": {
		Name: "Rock-It! Turret XL",
		Text: "Rock-It! Turrets attack up to 2 additional enemies for 60% damage.",
	},
	"TinkerMasteryTurretStorage": {
		Name: "Turret Storage",
		Text: "Increases Rock-It! Turret maximum charges from 2 to 3.",
	},
	"TinkerMasteryXTraLargeBombs": {
		Name: "X-Tra Large Bombs",
		Text: "Increases Xplodium Charge radius by 30%. Being stunned or rooted causes an Xplodium Charge to drop at Gazlowe's feet.\nThis effect has a 10 second cooldown.",
	},
	"TinkerSuperiorSchematics": {
		Name: "Superior Schematics",
		Text: "Gazlowe's Rock-It! Turrets have 25% increased range while he is within their attack range.",
	},
	"TinkerTalentFirinMahLazorz": {
		Name: "Firin' Mah Lazorz",
		Text: "Firing Deth Lazor causes all of Gazlowe's Rock-It! Turrets to fire a Deth Lazor of their own that deals 137 damage.",
	},
	"TracerBulletSprayMelee": {
		Name: "Bullet Spray",
		Text: "Increases Melee's radius by 30%, and causes it to damage all enemies in range.",
	},
	"TracerBulletTime": {
		Name: "Bullet Time",
		Text: "Basic Attacks lower the cooldown of Blink by 0.1 seconds.",
	},
	"TracerCompositionBPulseBomb": {
		Name: "Composition B",
		Text: "Successfully sticking a Pulse Bomb to an enemy Hero also drops another one at their feet that deals 50% damage and explodes slightly earlier.",
	},
	"TracerFocusFire": {
		Name: "Focus Fire",
		Text: "If an entire ammo magazine is unloaded on an enemy, the last bullet will deal 94.5 bonus damage.  This is equal to 35% of the total magazine.",
	},
	"TracerGetStuffedMelee": {
		Name: "Get Stuffed!",
		Text: "Reduces Melee's cooldown by 3 seconds. Hitting an enemy with Melee who is stuck with a Pulse Bomb causes the bomb to instantly explode and knocks the target away.",
	},
	"TracerIsthataHealthPack": {
		Name: "Is That a Health Pack?!",
		Text: "Increases Regeneration Globe and Healing Fountain healing by 100%.",
	},
	"TracerJumper": {
		Name: "Jumper",
		Text: "Increases Blink's charges by 1.",
	},
	"TracerLeechingRounds": {
		Name: "Leeching Rounds",
		Text: "Basic Attacks against Heroes heal Tracer for 20% of their damage dealt.",
	},
	"TracerLockedandLoaded": {
		Name: "Locked and Loaded",
		Text: "Reactivate Reload within the last 50% of its cast time to increase Tracer's Basic Attack damage by 40% for that magazine.",
	},
	"TracerPartingGift": {
		Name: "Parting Gift",
		Text: "Recall leaves behind 3 bombs that deal 240 damage each to different targets.",
	},
	"TracerPulseRounds": {
		Name: "Pulse Rounds",
		Text: "Increases Pulse Bomb's range and charge rate from Basic Attacks against Heroes by 100%.",
	},
	"TracerPulseStrikeMelee": {
		Name: "Pulse Strike",
		Text: "Increases Melee's Pulse Bomb charge from 10% to 20% against Heroes.",
	},
	"TracerQuantumSpike": {
		Name: "Quantum Spike",
		Text: "Pulse Bomb deals an additional 10% of the primary target's maximum Health.",
	},
	"TracerRicochetHeroWeapon": {
		Name: "Ricochet",
		Text: "Tracer's Basic Attacks have a 50% chance to hit another nearby enemy, prioritizing Heroes.",
	},
	"TracerSleightofHand": {
		Name: "Sleight of Hand",
		Text: "Reduces Reload time by 50%. This equals 20% more damage per second.",
	},
	"TracerSlipstreamRecall": {
		Name: "Slipstream",
		Text: "Increases the amount of time Recalled by 1 second.",
	},
	"TracerSpatialEcho": {
		Name: "Spatial Echo",
		Text: "Hero Takedowns grant 2 charges of Blink. ",
	},
	"TracerStickyBomb": {
		Name: "Sticky Bomb",
		Text: "Increases Pulse Bomb's radius by 50% and enemies hit are slowed by 70% for 3 seconds.",
	},
	"TracerTotalRecallRecall": {
		Name: "Total Recall",
		Text: "Recall's cooldown is increased by 8 seconds, but it heals Tracer equal to the amount of Health she lost during that time.",
	},
	"TracerTracerRounds": {
		Name: "Tracer Rounds",
		Text: "Basic Attacks reveal enemies for 4 seconds.",
	},
	"TracerUntouchable": {
		Name: "Untouchable",
		Text: "Quest: Takedowns increase Tracer's Basic Attack damage by 2%, up to 30%. These bonuses are lost on death.",
	},
	"TychusCombatTactician": {
		Name: "Combat Tactician",
		Text: "Basic Attacks reduce the cooldown of Run and Gun by 0.8 seconds.",
	},
	"TychusFullyLoaded": {
		Name: "Fully Loaded",
		Text: "Reduce Minigun's cooldown by 5 seconds.",
	},
	"TychusHeroicAbilityCommandeerOdin": {
		Name: "Commandeer Odin",
		Text: "Call down an Odin to pilot. The Odin deals increased Damage, has 100% increased Basic Attack range, and uses different Abilities. The Odin has 25 Armor and lasts 23 seconds.",
	},
	"TychusHeroicAbilityDrakkenLaserDrill": {
		Name: "Drakken Laser Drill",
		Text: "Call down a Laser Drill to attack nearby enemies, dealing 142 damage every second. Reactivate to assign a new target. Lasts 22 seconds.",
	},
	"TychusInTheRhythm": {
		Name: "In the Rhythm",
		Text: "Quest: While Minigun is active, Basic Attacks against Heroes permanently increase future Minigun durations by 0 seconds.",
	},
	"TychusMasterAssassin": {
		Name: "Master Assassin",
		Text: "Quest: Hero Takedowns increase Tychus's Attack Speed by 1%, up to 15%. \nReward: After getting 15 Takedowns, Tychus's Attack Speed is increased by an additional 10%.",
	},
	"TychusMasteryDrakkenLaserFocusingDiodes": {
		Name: "Focusing Diodes",
		Text: "Increases the range of the Drakken Laser Drill by 50%.  Deals increasing damage the longer it remains on a single target, up to 100% extra damage.",
	},
	"TychusMasteryFragGrenadeConcussionGrenade": {
		Name: "Concussion Grenade",
		Text: "Increases Frag Grenade's radius by 25% and the knockback distance by 100%.",
	},
	"TychusMasteryFragGrenadeTitanGrenade": {
		Name: "Titan Grenade",
		Text: "Frag Grenade deals bonus damage to Heroes equal to 5% of their maximum Health.",
	},
	"TychusMasteryLeadRain": {
		Name: "Lead Rain",
		Text: "Overkill also slows enemy movement speed by 20%.",
	},
	"TychusMasteryOdinBigRedButton": {
		Name: "Big Red Button",
		Text: "Odin lasts 50% longer and Ragnarok Missiles also launches a Nuclear Missile which lands 2.5 seconds later, dealing 319 damage in its area.",
	},
	"TychusMasteryOverkillArmorPiercingRounds": {
		Name: "Armor Piercing Rounds",
		Text: "Increases Overkill's damage to the primary target by 50%.",
	},
	"TychusMasteryQuarterback": {
		Name: "Quarterback",
		Text: "Increases the range of Frag Grenade by 50%.",
	},
	"TychusMasteryRunandGunDash": {
		Name: "Dash",
		Text: "Quest: Run and Gun increases Tychus's Movement Speed by 1% for each Regeneration Globe gathered, up to 25%. Lasts 2 seconds. \nReward: Once 25 Regeneration Globes have been gathered, the range of Run and Gun is increased by 50%.",
	},
	"TychusMasteryRunandGunStimPack": {
		Name: "Stim Pack",
		Text: "After using Run and Gun, gain 0% Attack Speed and 0% Movement Speed for 0 seconds",
	},
	"TychusMasteryShredderGrenade": {
		Name: "Shredder Grenade",
		Text: "Increases the explosion radius of Frag Grenade by 25%.",
	},
	"TychusMasterySprayNPray": {
		Name: "Spray 'n' Pray",
		Text: "Increases Overkill's range by 25%.",
	},
	"TychusNeosteelCoating": {
		Name: "Neosteel Coating",
		Text: "Activate to gain 75 Spell Armor for 3 seconds, reducing Ability Damage taken by 75%.",
	},
	"TychusPressTheAdvantage": {
		Name: "Press the Advantage",
		Text: "Run and Gun increases Tychus's Basic Attack range by 1.5 for 3 seconds.",
	},
	"TychusRelentlessSoldier": {
		Name: "Relentless Soldier",
		Text: "Being Stunned or Rooted grants 25 Armor for 3 seconds, reducing damage taken by 25%.",
	},
	"TychusRunAndGunBobAndWeave": {
		Name: "Bob and Weave",
		Text: "Run and Gun and Odin's Thrusters gain 2 additional charges and reduces Run and Gun's Mana cost from 50 to 25.",
	},
	"TychusSizzlinAttacks": {
		Name: "Sizzlin' Attacks",
		Text: "Basic Attacks deal bonus damage to Heroes equal to 1% of their maximum Health. Stacks with Minigun.",
	},
	"TychusThatsTheStuff": {
		Name: "That's the Stuff!",
		Text: "Minigun heals Tychus for 100% of the bonus damage dealt after it expires.",
	},
	"TychusTheBiggerTheyAre": {
		Name: "The Bigger They Are...",
		Text: "Increases Minigun damage bonus to 4% while enemy Heroes are above 35% Health, but Minigun no longer has any effect on targets below 35%.",
	},
	"TyraelHeroicAbilityJudgement": {
		Name: "Judgment",
		Text: "After 0.8 seconds, charge an enemy Hero dealing 150 damage and stunning them for 1.5 seconds. Nearby enemies are knocked away and take 75 damage.",
	},
	"TyraelHeroicAbilitySanctification": {
		Name: "Sanctification",
		Text: "After 0.5 seconds create a field of holy energy that makes allied Heroes Invulnerable. Lasts 3 seconds.",
	},
	"TyraelMasteryAngelicAbsorption": {
		Name: "Angelic Absorption",
		Text: "Enemies that attack Tyrael while shielded grant 240 Health over 4 seconds.",
	},
	"TyraelMasteryAngelicMight": {
		Name: "Angelic Might",
		Text: "Gain 80% increased damage on Tyrael's next Basic Attack within the next 5 seconds for each Hero hit by Smite.",
	},
	"TyraelMasteryElDruinsMightAngelsGrace": {
		Name: "Angel's Grace",
		Text: "After teleporting using El'druin's Might, gain 40% Movement Speed for 3 seconds.",
	},
	"TyraelMasteryElDruinsMightBladeOfJustice": {
		Name: "Blade of Justice",
		Text: "After teleporting using El'druin's Might, Tyrael's Attack Speed is increased by 50% for 5 seconds.",
	},
	"TyraelMasteryElDruinsMightHoradricReforging": {
		Name: "Horadric Reforging",
		Text: "If El'druin's Might hits an enemy, its cooldown is reduced by 5 seconds.",
	},
	"TyraelMasteryEvenInDeath": {
		Name: "Even In Death",
		Text: "Increases Archangel's Wrath's damage by 25%. Basic Abilities also can be used before exploding, but deal no damage.",
	},
	"TyraelMasteryHolyGround": {
		Name: "Holy Ground",
		Text: "Create a ring that blocks enemies from entering the area teleported to using El'druin's Might.",
	},
	"TyraelMasteryJudgmentAngelofJustice": {
		Name: "Angel of Justice",
		Text: "Increases the cast range of Judgment by 50%, and reduces the cooldown by 40 seconds.",
	},
	"TyraelMasteryPassiveProtectioninDeath": {
		Name: "Protection in Death",
		Text: "When Archangel's Wrath explodes, shield nearby allies for 50% of their max Health for 10 seconds.",
	},
	"TyraelMasteryPurgeEvil": {
		Name: "Purge Evil",
		Text: "Smite deals 30% more damage to Heroes.",
	},
	"TyraelMasteryRighteousnessReciprocate": {
		Name: "Reciprocate",
		Text: "When Righteousness expires, it explodes for 170 damage to nearby enemies.",
	},
	"TyraelMasterySalvation": {
		Name: "Salvation",
		Text: "Each allied Hero that is affected by Righteousness increases Tyrael's Shield amount by 45%.",
	},
	"TyraelMasterySanctificationHolyArena": {
		Name: "Holy Arena",
		Text: "Increases duration of Sanctification by 1 second and increases the damage of allies by 25%.",
	},
	"TyraelMasterySwiftRetribution": {
		Name: "Swift Retribution",
		Text: "Increases Smite's Movement Speed by 10% and duration by 1 second.",
	},
	"TyraelMasteryZealotry": {
		Name: "Zealotry",
		Text: "Increases the duration of Righteousness by 50%, and reduces its cooldown by 2 seconds.",
	},
	"TyrandeCelestialAttunement": {
		Name: "Celestial Attunement",
		Text: "Light of Elune removes Stuns and Silences from its target. When a Stun or Silence is removed this way, Light of Elune may be used for free within 4 seconds. This free cast cannot benefit from Celestial Attunement.",
	},
	"TyrandeDarnassianArchery": {
		Name: "Darnassian Archery",
		Text: "Basic Attacks against enemy Heroes grant a stacking 5% increased Attack Damage bonus for 4 seconds. This bonus is lost when Basic Attacking non-Heroes.",
	},
	"TyrandeElunesChosen": {
		Name: "Elune's Chosen",
		Text: "Activate to make Tyrande's Basic Attacks heal the target ally for 200% of the damage dealt. Lasts for 5 seconds.",
	},
	"TyrandeEmpower": {
		Name: "Empower",
		Text: "Reduce Sentinel's cooldown by 2 seconds. Every time it hits a Hero, reduce it's cooldown by an additional 4 seconds.",
	},
	"TyrandeEyesOfTheHuntress": {
		Name: "Eyes of the Huntress",
		Text: "After casting Shadowstalk, reveal all enemy Heroes for 3 seconds.",
	},
	"TyrandeHarshMoonlight": {
		Name: "Harsh Moonlight",
		Text: "Sentinel Slows the target by 25% and reduces their damage dealt by 25% for 4 seconds.",
	},
	"TyrandeHeroicAbilityShadowstalk": {
		Name: "Shadowstalk",
		Text: "Stealth and gain 30% Movement Speed. When Shadowstalk's Stealth is broken, gain 50% Attack Speed and retain its Movement Speed bonus for 5 seconds.",
	},
	"TyrandeHeroicAbilityStarfall": {
		Name: "Starfall",
		Text: "Deal 182 damage per second and slow enemies by 20% in an area. Lasts 6 seconds.",
	},
	"TyrandeHuntressFury": {
		Name: "Huntress' Fury",
		Text: "Increase the cast range of Hunter's Mark by 50%. Tyrande's Basic Attacks against targets with Hunter's Mark splash to nearby Heroes and Mercenaries.",
	},
	"TyrandeIcebladeArrows": {
		Name: "Iceblade Arrows",
		Text: "Increase Attack Speed by 25%. Basic Attacks against enemy Heroes reduce their damage dealt by 8% for 2 seconds. This effect stacks up to 5 times.",
	},
	"TyrandeKaldoreiResistance": {
		Name: "Kaldorei Resistance",
		Text: "Light of Elune grants the target 15 Spell Armor for 6 seconds. This effect stacks up to 3 times.",
	},
	"TyrandeMarkofMending": {
		Name: "Mark of Mending",
		Text: "Tyrande's Basic Attacks heal her for 1.8% of her maximum Health. Basic Attacks against targets with Hunter's Mark heal the attacker for 1.9% of their maximum Health.",
	},
	"TyrandeMasteryLightofEluneQuickeningBlessing": {
		Name: "Quickening Blessing",
		Text: "Light of Elune increases the target's Movement Speed by 30% for 3 seconds.",
	},
	"TyrandeMasteryLunarBlaze": {
		Name: "Lunar Blaze",
		Text: "Repeatable Quest: Hitting Heroes with Lunar Flare increases its damage by 3%.\nReward: After hitting 10 Heroes, Lunar Flare no longer has a Mana cost.\nReward: After hitting 20 Heroes, increase the range of Lunar Flare by 40%.",
	},
	"TyrandeMasteryLunarFlareShootingStar": {
		Name: "Shooting Star",
		Text: "Lunar Flare does 50% more damage. Hitting an enemy Hero will also refund the Mana cost.",
	},
	"TyrandeMasterySentinelPierce": {
		Name: "Pierce",
		Text: "Sentinel no longer stops at the first Hero hit, affecting all enemy Heroes along the path.",
	},
	"TyrandeMasteryShadowstalkHuntersSwiftness": {
		Name: "Hunter's Swiftness",
		Text: "Shadowstalk grants 30% Movement Speed for 5 seconds.",
	},
	"TyrandeMasteryStarfallCelestialWrath": {
		Name: "Celestial Wrath",
		Text: "Starfall applies Hunter's Mark to enemy Heroes while inside of its area of effect.",
	},
	"TyrandeMoonlitArrows": {
		Name: "Moonlit Arrows",
		Text: "Basic Attacks decrease the cooldown of Light of Elune by an additional 0.8 seconds.",
	},
	"TyrandeOverflowingLight": {
		Name: "Overflowing Light",
		Text: "While Tyrande is above 60% Health, Light of Elune's healing is increased by 30%.",
	},
	"TyrandeRanger": {
		Name: "Ranger",
		Text: "Sentinel's width is increased by 25% and deals more damage based on the distance traveled, up to 75%.\nRepeatable Quest: Hitting enemy Heroes with Sentinel increases the maximum damage bonus by 3%.",
	},
	"TyrandeRangersMark": {
		Name: "Ranger's Mark",
		Text: "Basic Attacks reduce the cooldown of Hunter's Mark by 1 second.\nRepeatable Quest: Every 50 Basic Attacks against enemy Heroes increase the duration of Hunter's Mark by 1 second.",
	},
	"TyrandeSearingArrows": {
		Name: "Searing Arrows",
		Text: "Activate to increase Basic Attack damage at the cost of Mana",
	},
	"TyrandeShootingStar": {
		Name: "Shooting Star",
		Text: "Increase Basic Attack range by 1.1. Every 10th Basic Attack against Heroes casts a free Lunar Flare at a random Hero near Tyrande's position.",
	},
	"TyrandeTrueshotAura": {
		Name: "Trueshot Aura",
		Text: "Gain Attack Damage aura",
	},
	"UtherBeaconOfLight": {
		Name: "Beacon of Light",
		Text: "While below 50% Health, Uther receives 200% more self-healing when healing others with Holy Light.",
	},
	"UtherCombatStyleHammerOfTheLightbringer": {
		Name: "Hammer of the Lightbringer",
		Text: "Quest: Basic Attack enemies 75 times.\nReward: Basic Attacks reduce the cooldown of Hammer of Justice by 1.5 seconds.\nPassive: Basic Attacks restore 1.5% of Uther's maximum Mana.",
	},
	"UtherEternalDevotionArmorOfFaith": {
		Name: "Armor of Faith",
		Text: "Being Stunned, Rooted, or Silenced causes Holy Light's cooldown to recharge 200% faster for 6 seconds.",
	},
	"UtherEternalDevotionDivineProtection": {
		Name: "Divine Protection",
		Text: "Applying Devotion to a Hero that is already affected by Devotion increases their Armor by an additional 15.",
	},
	"UtherEternalDevotionGuardianOfAncientKings": {
		Name: "Guardian of Ancient Kings",
		Text: "Healing a Hero that is Stunned, Rooted, or Silenced increases the Armor bonus they receive from Devotion from 15 to 50.",
	},
	"UtherHammerOfJusticePursuitOfJustice": {
		Name: "Pursuit of Justice",
		Text: "Hammer of Justice increases Uther's Movement Speed by 25% for 3 seconds.",
	},
	"UtherHammerOfJusticeWellMet": {
		Name: "Well Met",
		Text: "When used against Heroes, Hammer of Justice reduces the target's Movement Speed and damage dealt by 25% for 3 seconds.",
	},
	"UtherHandOfProtection": {
		Name: "Hand of Protection",
		Text: "Activate to make target ally Unstoppable for 1 second. Cannot be cast on Uther. Basic Attacks reduce Hand of Protection's cooldown by 5 seconds.",
	},
	"UtherHardenedFocus": {
		Name: "Hardened Focus",
		Text: "While above 80% life, your Basic Ability cooldowns regenerate 50% faster.",
	},
	"UtherHeroicAbilityDivineShield": {
		Name: "Divine Shield",
		Text: "Make an allied Hero Invulnerable and increase their Movement Speed by 20% for 3 seconds.",
	},
	"UtherHeroicAbilityDivineStorm": {
		Name: "Divine Storm",
		Text: "Deal 187 damage and Stun nearby enemies for 1.8 seconds.",
	},
	"UtherHolyFire": {
		Name: "Holy Fire",
		Text: "Deal 14 damage per second to nearby enemies. Basic Attacks against enemy Heroes increase this damage by 20% for 3 seconds. This can stack up to 3 times.",
	},
	"UtherHolyLightSilverTouch": {
		Name: "Silver Touch",
		Text: "Quest: Reduce damage from Heroic sources with Devotion.\nReward: After reducing damage 40 times, Holy Light's Mana cost is reduced from 90 to 70.\nReward: After reducing damage 80 times, Holy Light's Mana cost is reduced to 50, and its range is increased by 50%.",
	},
	"UtherHolyRadianceTyrsDeliverance": {
		Name: "Tyr's Deliverance",
		Text: "Allies healed by Holy Radiance receive 40% increased healing from all sources for 6 seconds.",
	},
	"UtherMasteryBenediction": {
		Name: "Benediction",
		Text: "Activate to reduce the Mana cost of Uther's next Basic Ability by 50 and its cooldown by 10 seconds.",
	},
	"UtherMasteryBlessedChampion": {
		Name: "Blessed Champion",
		Text: "For the next 5 seconds after using Holy Light, Uther's Basic Attacks heal him and nearby allies for 15% of the total amount healed by Holy Light.",
	},
	"UtherMasteryBulwarkOfLightDivineShield": {
		Name: "Bulwark of Light",
		Text: "Divine Shield lasts 2 seconds longer and its cooldown is reduced by 20 seconds.",
	},
	"UtherMasteryDivineHurricaneDivineStorm": {
		Name: "Divine Hurricane",
		Text: "Divine Storm's radius is increased by 50% and its cooldown is reduced by 20 seconds.",
	},
	"UtherMasteryHolyShock": {
		Name: "Holy Shock",
		Text: "Holy Light can be used on an enemy to do 50% of its healing amount as damage. When used this way, Uther receives its self-healing benefits, its cooldown is reduced by 6 seconds, and it refunds 45 Mana.",
	},
	"UtherMasteryRedemption": {
		Name: "Redemption",
		Text: "After Eternal Vanguard ends, Uther revives at the spirit's location with 50% of his maximum Health.\nThis effect has a 180 second cooldown.",
	},
	"UtherMasteryWaveofLightHolyRadiance": {
		Name: "Wave of Light",
		Text: "Quest: Damage or heal Heroes 60 times with Holy Radiance.\nReward: Increase the duration of Devotion by 0.5 seconds.\nPassive: Damaging or healing Heroes with Holy Radiance refunds 5 Mana and reduces its cooldown by 1 second.",
	},
	"ValeeraAmbushAssassinate": {
		Name: "Assassinate",
		Text: "Ambush deals 100% additional damage if no other enemy Heroes are within 4 range of the victim.",
	},
	"ValeeraAmbushDeathFromAbove": {
		Name: "Death From Above",
		Text: "Ambush now has 6 range and teleports Valeera behind the victim.",
	},
	"ValeeraBladeFlurryFatalFinesse": {
		Name: "Fatal Finesse",
		Text: "Quest: Each time Blade Flurry damages an enemy Hero, its damage increases by 3, up to 60.\nReward: After damaging 20 Heroes with Blade Flurry, permanently reduce its Energy cost by 20.",
	},
	"ValeeraCheapShotBlind": {
		Name: "Blind",
		Text: "When Cheap Shot expires, the victim is blinded for 2.5 seconds.",
	},
	"ValeeraCloakOfShadows": {
		Name: "Cloak of Shadows",
		Text: "Valeera is enveloped in a Cloak of Shadows, which immediately removes all damage over time effects from her. For 1.5 seconds, she becomes Unstoppable and gains 75 Spell Armor, reducing Ability Damage taken by 75%.\nUsing this Ability does not break Vanish.",
	},
	"ValeeraCloakOfShadowsEnvelopingShadows": {
		Name: "Enveloping Shadows",
		Text: "Whenever Valeera uses Vanish, she also gains Cloak of Shadows.",
	},
	"ValeeraColdBlood": {
		Name: "Cold Blood",
		Text: "Activate to increase the damage of Valeera's next Eviscerate by 100%.\nActivating Cold Blood does not break Vanish.",
	},
	"ValeeraCombatReadiness": {
		Name: "Combat Readiness",
		Text: "Each Combo Point spent grants 50 Physical Armor against the next enemy Hero Basic Attack, reducing its damage by 50%. Stores up to 3 charges.",
	},
	"ValeeraCripplingPoison": {
		Name: "Crippling Poison",
		Text: "Activate to cause Valeera's next instance of Spell Damage to slow enemy Movement Speed by 25% for 5 seconds.\nActivating Crippling Poison does not break Vanish.",
	},
	"ValeeraEviscerateExposeArmor": {
		Name: "Expose Armor",
		Text: "Eviscerating an enemy with 3 Combo Points causes them to lose 25 Armor for 2 seconds, increasing all damage taken by 25%.",
	},
	"ValeeraEviscerateSliceAndDice": {
		Name: "Slice and Dice",
		Text: "Eviscerating an enemy with 3 Combo Points causes Valeera's next 3 Basic Attacks to have their Attack Speed increased by 150%. Lasts 3 seconds.",
	},
	"ValeeraGarroteHemorrhage": {
		Name: "Hemorrhage",
		Text: "Valeera's Basic Attacks deal 25% additional damage to enemies affected by Garrote.",
	},
	"ValeeraGarroteRupture": {
		Name: "Rupture",
		Text: "Valeera's Basic Attacks increase the damage over time effect of Garrote by 5%, stacking up to 30%, and refresh the duration.",
	},
	"ValeeraGarroteStrangle": {
		Name: "Strangle",
		Text: "Using Garrote on a Hero reduces their Spell Power by 25% for 7 seconds.",
	},
	"ValeeraSinisterStrikeMutilate": {
		Name: "Mutilate",
		Text: "Sinister Strike now hits with both blades, dealing 100% additional damage, but its range is reduced by 1.",
	},
	"ValeeraSinisterStrikeRelentlessStrikes": {
		Name: "Relentless Strikes",
		Text: "Reduces the Energy cost of Sinister Strike by 10.",
	},
	"ValeeraSinisterStrikeSealFate": {
		Name: "Seal Fate",
		Text: "Sinister Strike deals 50% additional damage and generates an additional Combo Point when used against silenced, rooted, or stunned enemy Heroes.",
	},
	"ValeeraSmokeBomb": {
		Name: "Smoke Bomb",
		Text: "Create a cloud of smoke. While in the smoke, Valeera is Unrevealable, can pass through other units, and gains 25 Armor, reducing damage taken by 25%. Valeera can continue to attack and use abilities without being revealed. Lasts 5 seconds.\nUsing this Ability does not break Vanish.",
	},
	"ValeeraSmokeBombAdrenalineRush": {
		Name: "Adrenaline Rush",
		Text: "While in the Smoke Bomb, Valeera regenerates an additional 30 Energy per second.",
	},
	"ValeeraThistleTea": {
		Name: "Thistle Tea",
		Text: "Activate to instantly restore 100 Energy.\nActivating Thistle Tea does not break Vanish.",
	},
	"ValeeraVanishElusiveness": {
		Name: "Elusiveness",
		Text: "Increases Valeera's Movement Speed while Vanished by an additional 10%.",
	},
	"ValeeraVanishInitiative": {
		Name: "Initiative",
		Text: "Ambush, Cheap Shot, and Garrote award 1 Combo Points.",
	},
	"ValeeraVanishNightslayer": {
		Name: "Nightslayer",
		Text: "Reduces the cooldown of Vanish by 3 seconds.",
	},
	"ValeeraVanishSubtlety": {
		Name: "Subtlety",
		Text: "After remaining Vanished for at least 4 seconds, using Ambush, Cheap Shot, or Garrote causes Valeera to regenerate an additional 10 Energy per second for 5 seconds.",
	},
	"ValeeraVigor": {
		Name: "Vigor",
		Text: "Valeera regenerates an additional 2 Energy per second.\nReward: After Gathering 20 Regen Globes, increase Valeera's maximum Energy to 120.",
	},
	"ValeeraWoundPoison": {
		Name: "Wound Poison",
		Text: "Activate to cause Valeera's next instance of Spell Damage to also reduce healing received by the victim by 50% for 5 seconds.\nActivating Wound Poison does not break Vanish.",
	},
	"VarianBannerOfDalaran": {
		Name: "Banner of Dalaran",
		Text: "Activate to place a Banner that grants 20% increased Spell Power to nearby allied Heroes. Lasts 12 seconds.",
	},
	"VarianBannerOfIronforge": {
		Name: "Banner of Ironforge",
		Text: "Activate to place a Banner that grants 20 Armor to nearby allied Heroes, reducing damage taken by 20%. Lasts 12 seconds.",
	},
	"VarianBannerOfStormwind": {
		Name: "Banner of Stormwind",
		Text: "Activate to place a Banner that grants 25% increased Movement Speed to nearby allied Heroes. Lasts 12 seconds.",
	},
	"VarianBannersGloryToTheAlliance": {
		Name: "Glory to the Alliance",
		Text: "Banner now also increases health regeneration and all healing received for nearby allied Heroes by 50%, and the cooldown is reduced by 20 seconds.",
	},
	"VarianChargeJuggernaut": {
		Name: "Juggernaut",
		Text: "Charge deals bonus damage to Heroes equal to 4% of their maximum Health.",
	},
	"VarianChargeWarbringer": {
		Name: "Warbringer",
		Text: "Reduces Charge cooldown by 8 seconds and Mana cost by 50%.",
	},
	"VarianColossusSmash": {
		Name: "Colossus Smash",
		Text: "Smash a target enemy, dealing 160 damage and lowering their Armor by 25 for 3 seconds, causing them to take 25% increased damage.\nPassive: Base Attack Damage increased by 100%.\nPassive: Maximum Health reduced by 10%.",
	},
	"VarianColossusSmashMasterAtArms": {
		Name: "Master at Arms",
		Text: "Colossus Smash affects all enemies near the target, and its cooldown is reduced by 10 seconds.",
	},
	"VarianDemoralizingShout": {
		Name: "Demoralizing Shout",
		Text: "Activate to demoralize nearby enemy Heroes, reducing damage they deal by 25% for 5 seconds.",
	},
	"VarianHighKingsQuestQuest": {
		Name: "High King's Quest",
		Text: "Quest: Hit 50 Heroes with Basic Attacks.\nQuest: Participate in 5 Hero Takedowns.\nQuest: Gather 20 Regeneration Globes.\nReward: Completing a Quest grants 10 Base Attack Damage. Completing all 3 Quests grants an additional 45 Base Attack Damage.",
	},
	"VarianLionsFangLionsMawQuest": {
		Name: "Lion's Maw",
		Text: "Quest: Every time Lion's Fang hits a Hero, increase its damage by 4, up to 120.\nReward: After hitting 30 Heroes, the slow is increased to 50% and its duration is increased to 2 seconds.",
	},
	"VarianMortalStrike": {
		Name: "Mortal Strike",
		Text: "Heroes hit by Heroic Strike receive 40% reduced healing for 4 seconds.",
	},
	"VarianParryLivebytheSword": {
		Name: "Live by the Sword",
		Text: "Increase the duration of Parry by 40%. If at least 4 Hero Basic Attacks are blocked with a single Parry, its cooldown is reduced by 2 seconds.",
	},
	"VarianParryOverpower": {
		Name: "Overpower",
		Text: "When Parry blocks a Hero's Basic Attack, Heroic Strike's cooldown is refreshed and the next one does 20% more damage.",
	},
	"VarianParryShieldWall": {
		Name: "Shield Wall",
		Text: "Parry now grants Protected, preventing all incoming damage for the duration.",
	},
	"VarianSecondWind": {
		Name: "Second Wind",
		Text: "Basic Attacks heal Varian for 1.2% of his maximum Health. While below 50% Health, they also heal him for 50% of the damage dealt.",
	},
	"VarianShatteringThrow": {
		Name: "Shattering Throw",
		Text: "Activate to throw a sword at a target Hero that deals 50 damage, and up to 1400 bonus damage to their Shields.\nPassive: Basic Attacks against Heroes deal up to 200% bonus damage to Shields.",
	},
	"VarianTaunt": {
		Name: "Taunt",
		Text: "Silence a target Hero and force them to attack Varian for 1.2 seconds.\nPassive: Maximum Health increased by 30%.\nPassive: Gain 15 Armor.",
	},
	"VarianTauntVigilance": {
		Name: "Vigilance",
		Text: "Being hit by a Hero Basic Attack reduces the cooldown of Taunt by 1 seconds.",
	},
	"VarianTwinBladesOfFury": {
		Name: "Twin Blades of Fury",
		Text: "Basic Attacks reduce Heroic Strike's cooldown by 9 seconds, and increase Varian's Movement Speed by 30% for 2 seconds.\nPassive: Attack Speed increased by 100%.\nPassive: Base Attack Damage reduced by 20%.",
	},
	"VarianTwinBladesOfFuryFrenzy": {
		Name: "Frenzy",
		Text: "Twin Blades of Fury increases Varian's Attack Speed by an additional 15%, and Basic Attacks increase his Movement Speed by an additional 10%.",
	},
	"VarianVictoryRush": {
		Name: "Victory Rush",
		Text: "Every 60 seconds, Varian's next Basic Attack will heal him for 400 Health. When a nearby enemy Minion or Monster dies, the cooldown is reduced by 15 seconds.",
	},
	"WitchDoctorAnnihilatingSpirits": {
		Name: "Annihilating Spirit",
		Text: "Increases the range of Ravenous Spirit by 50% and Movement Speed by 30%.",
	},
	"WitchDoctorBigVoodoo": {
		Name: "Big Voodoo",
		Text: "Increases the Health and Mana bonuses from Voodoo Ritual by 100%.",
	},
	"WitchDoctorBloodRitual": {
		Name: "Blood Ritual",
		Text: "If an enemy dies while poisoned by Voodoo Ritual, restore 1.9% of Nazeebo's maximum Health and Mana.",
	},
	"WitchDoctorDeadRush": {
		Name: "Dead Rush",
		Text: "Zombie Wall deals 75% more damage. When it expires up to 5 remaining Zombies uproot and attack nearby enemies for 3 seconds.",
	},
	"WitchDoctorFreshCorpses": {
		Name: "Fresh Corpses",
		Text: "Zombie Wall cooldown reduced by 0 seconds.",
	},
	"WitchDoctorGuardianToads": {
		Name: "Guardian Toads",
		Text: "Hitting an enemy Hero with Plague of Toads gives 25 Armor for 2 seconds, reducing damage taken by 25%.",
	},
	"WitchDoctorHeroicAbilityGargantuan": {
		Name: "Gargantuan",
		Text: "Summon a Gargantuan to guard an area for 20 seconds. Deals 100 damage when summoned, attacks for 150 damage, and can be ordered to stomp nearby enemies.",
	},
	"WitchDoctorHeroicAbilityRavenousSpirits": {
		Name: "Ravenous Spirit",
		Text: "Channel a Ravenous Spirit that deals 206.8 damage per second. Cannot move while channeling. Lasts for 8 seconds.",
	},
	"WitchDoctorHexedCrawlers": {
		Name: "Hexed Crawlers",
		Text: "Corpse Spiders restore 0.8% of Nazeebo's maximum Health and Mana when they attack an enemy Hero.",
	},
	"WitchDoctorHumongoid": {
		Name: "Humongoid",
		Text: "Reduce Gargantuan's cooldown by 40 seconds and its Mana cost by 50%.",
	},
	"WitchDoctorPandemic": {
		Name: "Pandemic",
		Text: "Reward: After hitting 40 Heroes with Plague of Toads, it spawns 2 additional toads.",
	},
	"WitchDoctorRingOfPoison": {
		Name: "Ring of Poison",
		Text: "Zombie Wall lasts 1 second longer, and the center is filled with poison that deals a total of 318.8 damage over 4 seconds. This damage starts small and increases over the duration.",
	},
	"WitchDoctorSoulHarvest": {
		Name: "Soul Harvest",
		Text: "Activate to increase Nazeebo's Health and Spell Power by 7% for each nearby enemy, up to a maximum of 35%. Lasts 15 seconds.",
	},
	"WitchDoctorSpiderColony": {
		Name: "Spider Colony",
		Text: "Corpse Spider Attacks against Heroes reduce the cooldowns of Zombie Wall and Plague of Toads by 0.2 seconds.",
	},
	"WitchDoctorSpiritofArachyr": {
		Name: "Spirit of Arachyr",
		Text: "If Corpse Spiders hits only one enemy, it creates 2 additional spiders.",
	},
	"WitchDoctorSuperstition": {
		Name: "Superstition",
		Text: "Gain 40 Spell Armor. Heroic Basic Attacks against Nazeebo remove this bonus for 3 seconds.",
	},
	"WitchDoctorThingOfTheDeep": {
		Name: "Thing of the Deep",
		Text: "Increases the range of Nazeebo's Basic Abilities by 20%.",
	},
	"WitchDoctorToadAffinity": {
		Name: "Toad Affinity",
		Text: "Damage against heroes reduces cooldown",
	},
	"WitchDoctorToadsofHugeness": {
		Name: "Toads of Hugeness",
		Text: "Increase the damage and explosion radius of Plague of Toads by 25% after each hop, up to a maximum of 100%.",
	},
	"WitchDoctorVileInfection": {
		Name: "Vile Infection",
		Text: "Quest: Reach 175 stacks of Voodoo Ritual.\nReward: After reaching 175 stacks of Voodoo Ritual, it can also be applied to Heroes and its damage is increased by 200%.",
	},
	"WitchDoctorWidowmakers": {
		Name: "Widowmakers",
		Text: "Reward: After Corpse Spiders attack Heroes 100 times, their attack damage is increased by 30%.",
	},
	"WizardAetherWalker": {
		Name: "Aether Walker",
		Text: "If Li-Ming hasn't taken damage in the last 5 seconds, Teleport costs no mana and its cooldown is decreased by 2 seconds.",
	},
	"WizardArcaneOrbArcaneOrbit": {
		Name: "Arcane Orbit",
		Text: "Arcane Orb travels 25% farther, doing up to 25% more damage. ",
	},
	"WizardArcaneOrbTriumvirate": {
		Name: "Triumvirate",
		Text: "If Arcane Orb hits an enemy Hero after traveling at least 50% of its base range, the cooldown is reduced by 50%.",
	},
	"WizardArcaneOrbZeisVengeance": {
		Name: "Zei's Vengeance",
		Text: "Arcane Orb does up to 25% more damage to enemies based on distance. Additionally, its cooldown is reduced by 2 seconds.",
	},
	"WizardArchonPurePower": {
		Name: "Archon: Pure Power",
		Text: "Activate to enter Archon form, allowing Li-Ming to repeatedly use Disintegrate, but be unable to cast other Abilities.",
	},
	"WizardAstralPresence": {
		Name: "Astral Presence",
		Text: "Li-Ming's Mana regeneration is increased by 100% while below 25% Mana.",
	},
	"WizardCannoneer": {
		Name: "Cannoneer",
		Text: "When Li-Ming uses an Ability, her next Basic Attack's damage is increased by 75% and deals Spell damage instead of Physical.  Stacks up to 3 times.",
	},
	"WizardDisintegrateTemporalFlux": {
		Name: "Temporal Flux",
		Text: "Disintegrate gradually slows enemies caught in its beam, to a maximum of 60%.",
	},
	"WizardDominance": {
		Name: "Dominance",
		Text: "Takedowns restore 19.9% of Li-Ming's maximum Health.",
	},
	"WizardFireflies": {
		Name: "Fireflies",
		Text: "Drastically increases Magic Missiles speed. Its cooldown is reduced by 1 second, and its Mana cost is reduced by 5.",
	},
	"WizardForceArmor": {
		Name: "Force Armor",
		Text: "When Magic Missiles damages an enemy Hero, gain 50 Spell Armor against the next enemy Ability, reducing the damage taken by 50%. Gain 1 charge per cast.\nStores up to 4 charges.",
	},
	"WizardGlassCannon": {
		Name: "Glass Cannon",
		Text: "Increase Li-Ming's Spell Power by 15%, but decrease her maximum Health by 15%.",
	},
	"WizardHeroicAbilityDisintegrate": {
		Name: "Disintegrate",
		Text: "Channel a powerful beam, dealing 440 damage over 2.5 seconds to enemies while they are in it. The direction of the beam changes with your mouse cursor position.",
	},
	"WizardHeroicAbilityWaveOfForce": {
		Name: "Wave Of Force",
		Text: "Knock away all enemies from an area and deal 160 damage.",
	},
	"WizardMagicMissilesChargedBlast": {
		Name: "Charged Blast",
		Text: "Basic Attacking a target recently hit by a Magic Missile does an extra 87 damage.",
	},
	"WizardMagicMissilesMirrorball": {
		Name: "Mirrorball",
		Text: "Magic Missiles fires an additional 2 missiles, but its Mana cost is increased by 5.",
	},
	"WizardMagicMissilesSeeker": {
		Name: "Seeker",
		Text: "If three Magic Missiles hit the same target, the third one deals an additional 125 damage.",
	},
	"WizardPowerHungry": {
		Name: "Power Hungry",
		Text: "Regeneration Globes restore 100% more Mana and grant 10% Spell Power for 20 seconds.",
	},
	"WizardTalRashasElements": {
		Name: "Tal Rasha's Elements",
		Text: "Using Abilities grants 5% Spell Power, to a maximum of 20%.  This bonus is reset when the same Ability is used within a chain.",
	},
	"WizardTeleportCalamity": {
		Name: "Calamity",
		Text: "Teleport does 300 damage to enemy Heroes near the destination.",
	},
	"WizardTeleportDiamondSkin": {
		Name: "Diamond Skin",
		Text: "Teleport grants a Shield equal to 20% of Li-Ming's maximum Health for 4 seconds.",
	},
	"WizardTeleportIllusionist": {
		Name: "Illusionist",
		Text: "If Li-Ming loses more than 20% of her Health at once, its cooldown is instantly refreshed. \nThis effect has a 4 second cooldown.\nPassive: Increases Teleport range by 50%",
	},
	"WizardWaveOfForceRepulsion": {
		Name: "Repulsion",
		Text: "Increases Wave of Force knockback distance by 150% and increases its cast range by 100%.",
	},
	"ZagaraCombatStyleMedusaBlades": {
		Name: "Medusa Blades",
		Text: "Basic Attacks deal 33% damage to three nearby targets.",
	},
	"ZagaraHeroicAbilityDevouringMaw": {
		Name: "Devouring Maw",
		Text: "Summon a Devouring Maw that devours enemies for 4 seconds. Devoured enemies cannot fight and take 80 damage per second.\nUsable on Unstoppable enemies.",
	},
	"ZagaraHeroicAbilityNydusAssault": {
		Name: "Nydus Network",
		Text: "Summon a Nydus Worm on Creep anywhere that Zagara has vision. Zagara can enter a Nydus Worm and travel to any other Nydus Worm by right-clicking near it. While inside a Nydus Worm, Zagara regenerates 9.8% Health and Mana per second.\nStores up to 2 charges. Maximum of 10 Nydus Worms at a time.\nPassive: Creep spreads 15% farther. \nPassive: While on Creep, each Basic Attack reduces all of Zagara's cooldowns by .75 seconds.",
	},
	"ZagaraMasteryBanelingMassacre": {
		Name: "Baneling Massacre",
		Text: "Gain 2 additional charges of Banelings.",
	},
	"ZagaraMasteryBileDrop": {
		Name: "Bile Drop",
		Text: "Quest: Passively increases the damage of Infested Drop by 50%. Each Hero hit by the impact damage of Infested Drop increases this bonus by 5%.\nReward: Upon reaching 150% bonus damage, increase the impact radius of Infested Drop by 20%.",
	},
	"ZagaraMasteryBroodExpansion": {
		Name: "Brood Expansion",
		Text: "Reduces the cooldown of Hunter Killer by 6 seconds.",
	},
	"ZagaraMasteryCentrifugalHooks": {
		Name: "Centrifugal Hooks",
		Text: "Banelings can travel twice as far before exploding.",
	},
	"ZagaraMasteryCorpseFeeders": {
		Name: "Corpse Feeders",
		Text: "Reduces the cooldown of Infested Drop by 3 seconds and Roachlings take 40% less damage from non-Heroic sources.",
	},
	"ZagaraMasteryCorrosiveSaliva": {
		Name: "Corrosive Saliva",
		Text: "When used against Heroes, Hunter Killers deal additional damage equal to 2% of their maximum Health.",
	},
	"ZagaraMasteryEndlessCreep": {
		Name: "Endless Creep",
		Text: "The cast range for Creep Tumor is increased by 2000%. Creep Tumors now last 600 seconds. While on Creep, Zagara gains an additional 10% Movement Speed.",
	},
	"ZagaraMasteryEnvenomedSpines": {
		Name: "Envenomed Spines",
		Text: "Activate to have Zagara's next Basic Attack apply 230 damage over 5 seconds.",
	},
	"ZagaraMasteryGroovedSpines": {
		Name: "Grooved Spines",
		Text: "Hunter Killer has its range increased by 35% and damage increased by 20%.",
	},
	"ZagaraMasteryHydraliskTransfusion": {
		Name: "Hydralisk Transfusion",
		Text: "Zagara is healed for 75% of the damage dealt by Hunter Killers' Basic Attacks to Heroes.",
	},
	"ZagaraMasteryInfest": {
		Name: "Infest",
		Text: "Nearby Ranged Minions deal an additional 100% damage, plus an additional 1% per 1000 Siege damage Zagara has dealt. Can be toggled on or off.",
	},
	"ZagaraMasteryMetabolicBoost": {
		Name: "Metabolic Boost",
		Text: "Movement Speed boost on Creep increased to 40%.",
	},
	"ZagaraMasteryMutalisk": {
		Name: "Mutalisk",
		Text: "Hunter Killer now spawns a Mutalisk. Mutalisks have a bounce attack and last for 30 seconds.",
	},
	"ZagaraMasteryProtectiveCoating": {
		Name: "Protective Coating",
		Text: "While on Creep, Zagara gains 20 Armor, taking 20% less damage.",
	},
	"ZagaraMasteryReconstitution": {
		Name: "Reconstitution",
		Text: "Health Restoration bonus on Creep increased by 200%.",
	},
	"ZagaraMasterySerratedSpines": {
		Name: "Serrated Spines",
		Text: "Quest: Each Basic Attack against a Hero increases Zagara's Attack Damage by 0.3.",
	},
	"ZagaraMasteryTumorClutch": {
		Name: "Tumor Clutch",
		Text: "Creep Tumor Mana cost removed. Cooldown decreased to 10 seconds.",
	},
	"ZagaraMasteryTyrantMaw": {
		Name: "Tyrant Maw",
		Text: "Devouring Maw deals 50% more damage and getting a Takedown reduce its cooldown by 25 seconds.",
	},
	"ZagaraMasteryVolatileAcid": {
		Name: "Volatile Acid",
		Text: "Banelings can travel 50% further before exploding and their damage to non-Heroes is increased by 50%.",
	},
	"ZagaraViscousAcid": {
		Name: "Viscous Acid",
		Text: "Banelings slow enemies by 20% for 1.5 seconds.",
	},
	"ZaryaDefensiveShielding": {
		Name: "Defensive Shielding",
		Text: "Upon expiration or breaking, Personal Barrier and Shield Ally grant 75 Physical Armor against the next 2 Hero Basic Attacks for 6 seconds, reducing the damage taken by 75%.",
	},
	"ZaryaEnergyBornInBattle": {
		Name: "Born in Battle",
		Text: "While at or above 75 Energy, Zarya's cooldowns regenerate 25% faster.",
	},
	"ZaryaEnergyEnduranceTraining": {
		Name: "Endurance Training",
		Text: "While at or above 75 Energy, Zarya gains 20 Armor, reducing all damage taken by 20%.",
	},
	"ZaryaEnergyHitMe": {
		Name: "Hit Me",
		Text: "Damage absorbed by Zarya's Shields contributes 15% more Energy.",
	},
	"ZaryaEnergyMaximumChargeQuest": {
		Name: "Maximum Charge",
		Text: "Regen Globes now instantly grant 20 Energy.\nQuest: Spend 150 seconds at or above 50 Energy.\nReward: Gain an additional 20 maximum Energy.",
	},
	"ZaryaExpulsionZoneClearOut": {
		Name: "Clear Out",
		Text: "Increases Expulsion Zone's radius by 15% and Zarya's maximum Energy temporarily increases by 15 per enemy Hero hit for 10 seconds.",
	},
	"ZaryaGravitonSurgeGravityKills": {
		Name: "Gravity Kills",
		Text: "Increases Graviton Surge's duration by 1 second and causes it to deal 175 damage over the duration.",
	},
	"ZaryaHeroicAbilityExpulsionZone": {
		Name: "Expulsion Zone",
		Text: "Launch a gravity bomb that deals 124 damage and creates an expulsion zone for 3.5 seconds. Enemies who enter the affected area are knocked back and have their Movement Speed reduced by 50% for 1 second.",
	},
	"ZaryaHeroicAbilityGravitonSurge": {
		Name: "Graviton Surge",
		Text: "Launch a gravity bomb that detonates after 1 second and draws enemy Heroes toward the center for 2.5 seconds.",
	},
	"ZaryaPainIsTemporary": {
		Name: "Pain is Temporary",
		Text: "Activate to consume all Energy and gain a Shield that absorbs 0.5% of Zarya's maximum Health per Energy consumed and lasts for 3 seconds. Zarya may only have one personal Shield active on herself at a time.",
	},
	"ZaryaParticleGrenadeDemolitionsExpertQuest": {
		Name: "Demolitions Expert",
		Text: "Quest: For every 5 enemy Heroes hit by Particle Grenade, its recharge rate lowers by 0.4 seconds, up to 3.5 seconds.\nReward: Once Particle Grenade has hit 40 enemy Heroes, its radius is increased by 15%.",
	},
	"ZaryaParticleGrenadeGrenadier": {
		Name: "Grenadier",
		Text: "All Particle Grenade charges are returned at once.",
	},
	"ZaryaParticleGrenadePinpointAccuracy": {
		Name: "Pinpoint Accuracy",
		Text: "Particle Grenade deals 28.4 more damage to enemies hit by the center of the blast.",
	},
	"ZaryaParticleGrenadePlasmaShock": {
		Name: "Plasma Shock",
		Text: "Particle Grenade slows Movement Speed by 35% for 2 seconds to enemies near the center.",
	},
	"ZaryaPersonalBarrierExplosiveBarrier": {
		Name: "Explosive Barrier",
		Text: "Upon expiration or breaking, Personal Barrier explodes, dealing 100 damage to nearby enemies.",
	},
	"ZaryaPersonalBarrierIAmTheStrongest": {
		Name: "I Am the Strongest",
		Text: "Personal Barrier absorbs an additional 140 damage.",
	},
	"ZaryaPersonalBarrierSpellBarrier": {
		Name: "Spell Barrier",
		Text: "Upon expiration or breaking, Personal Barrier grants Zarya 75 Spell Armor for 3 Seconds, reducing the damage taken from Abilities by 75%.",
	},
	"ZaryaPersonalBarrierUnstoppableCompetitor": {
		Name: "Unstoppable Competitor",
		Text: "Personal Barrier makes Zarya Unstoppable for up to 2 seconds.",
	},
	"ZaryaShieldAllyCleansingShield": {
		Name: "Cleansing Shield",
		Text: "Causes Shield Ally to remove all disabling effects and reduces its cooldown by 3 seconds.",
	},
	"ZaryaShieldAllyGainTrain": {
		Name: "Gain Train",
		Text: "Shield Ally now grants an untalented Shield to a nearby ally upon impact.",
	},
	"ZaryaShieldAllyGiveMeTwentyQuest": {
		Name: "Give Me Twenty",
		Text: "Quest: Regeneration Globes increase Shield Ally absorb amount by 15, up to 300 permanently. \nReward: After collecting 20 Regen Globes, permanently reduce the cooldown of Shield Ally by 2 seconds.",
	},
	"ZaryaShieldAllySpeedBarrier": {
		Name: "Speed Barrier",
		Text: "Shield Ally increases the allied Hero's Movement Speed by 50% for the duration.",
	},
	"ZaryaShieldAllyTogetherWeAreStrong": {
		Name: "Together We Are Strong",
		Text: "Every 6.2 damage done by allies while under Shield Ally contributes 1 Energy, up to 30 Energy per Shield.",
	},
	"ZaryaUnyieldingDefender": {
		Name: "Unyielding Defender",
		Text: "Activate to reset the cooldown of Personal Barrier and Shield Ally.",
	},
	"ZaryaWeaponFeelTheHeat": {
		Name: "Feel the Heat",
		Text: "Zarya's Basic Attack deals 50% additional damage to enemies in melee range.",
	},
	"ZaryaWeaponToTheLimit": {
		Name: "To the Limit",
		Text: "While above 50 Energy, Zarya's Basic Attack size is increased by 35%.",
	},
	"ZeratulComboSlash": {
		Name: "Combo Slash",
		Text: "After using an Ability, Zeratul's next Basic Attack within 6 seconds deals 40% additional damage.",
	},
	"ZeratulGiftoftheXelNaga": {
		Name: "Gift of the Xel’Naga",
		Text: "Allies are no longer affected by Void Prison, and enemies are slowed by 50% for 3 seconds once Void Prison ends.",
	},
	"ZeratulGrimTask": {
		Name: "Grim Task",
		Text: "Quest: Hero Takedowns increase Spell Power by 5%, up to 50% . This bonus Spell Power is lost on death.",
	},
	"ZeratulHeroicAbilityShadowAssault": {
		Name: "Shadow Assault",
		Text: "Basic Attacks cause Zeratul to charge at enemies and have 20% increased Attack Speed. Lasts for 4 seconds.",
	},
	"ZeratulHeroicAbilityVoidPrison": {
		Name: "Void Prison",
		Text: "Slows time in an area to a near standstill, placing allies and enemies in Time Stop for 5 seconds. Zeratul is not affected.",
	},
	"ZeratulMasterWarpBlade": {
		Name: "Master Warp-Blade",
		Text: "Every third consecutive Basic Attack against the same target deals 125% bonus damage.",
	},
	"ZeratulMasteryGreaterCleaveCleave": {
		Name: "Greater Cleave",
		Text: "Increases the radius of Cleave by 33%.",
	},
	"ZeratulMendingStrikes": {
		Name: "Mending Strikes",
		Text: "Basic Attacks heal for 35% of the damage dealt.",
	},
	"ZeratulNerazimFury": {
		Name: "Nerazim Fury",
		Text: "Shadow Assault grants 30% Life Steal, and the duration is increased by 50%.",
	},
	"ZeratulRendingCleave": {
		Name: "Rending Cleave",
		Text: "Cleave deals an additional 40% damage over 3 seconds.",
	},
	"ZeratulSeekerInTheDark": {
		Name: "Seeker in the Dark",
		Text: "Singularity Spike takes 50% longer to explode. It can be reactivated to teleport to the target, granting 30% increased Move Speed for 3 seconds.",
	},
	"ZeratulSentencedtoDeath": {
		Name: "Sentenced to Death",
		Text: "Deal 40% increased damage to enemies while they have a Singularity Spike attached to them.",
	},
	"ZeratulShadowHunter": {
		Name: "Shadow Hunter",
		Text: "Quest: Gather Regeneration Globes to lower the Mana cost of Blink by 2 .\nReward: Upon gathering 20 Regeneration Globes, Basic Attacks reduce the cooldown of Blink by 1 second.",
	},
	"ZeratulShroudofAdun": {
		Name: "Shroud of Adun",
		Text: "Zeratul gains a shield equal to 15% of his Maximum Health over 5 seconds while under Permanent Cloak.",
	},
	"ZeratulSlipintoShadow": {
		Name: "Slip into Shadow",
		Text: "Blink gains an additional charge, but its cooldown is increased by 8 seconds.",
	},
	"ZeratulVoidSlash": {
		Name: "Void Slash",
		Text: "If Cleave hits more than one enemy Hero, it deals 40% increased damage and its cooldown is reduced by 3 seconds.",
	},
	"ZeratulVorpalBlade": {
		Name: "Vorpal Blade",
		Text: "Activate to teleport to Zeratul's last non-structure Basic Attack target within 3 seconds. The target is revealed during these 3 seconds.",
	},
	"ZeratulWormhole": {
		Name: "Wormhole",
		Text: "For 2.1 seconds, reactivate Blink to return to the point where it was cast from.",
	},
	"ZuljinAmaniRageTalent": {
		Name: "Amani Rage",
		Text: "Activate to cause Zul'jin to instantly lose 50% of his current Health and heal for that amount over 6 seconds. Zul'jin gains 10 Armor while regenerating health from Amani Rage.",
	},
	"ZuljinAmaniResilience": {
		Name: "Amani Resilience",
		Text: "Increase Taz'dingo's duration by 1 second. Upon expiring, Zul'jin heals for 50% of the damage he dealt during Taz'dingo.",
	},
	"ZuljinArcaniteAxes": {
		Name: "Arcanite Axes",
		Text: "Reduce Twin Cleave's cooldown by 2 seconds.\nQuest: Twin Cleave's damage is permanently increased by 0.8 every time it hits a Hero.",
	},
	"ZuljinBoneslicer": {
		Name: "Boneslicer",
		Text: "Grievous Throw now pierces through all enemies hit and restores 15 Mana per Hero hit. Additionally, Grievous Throw's mark is no longer removed by Basic Attacks.",
	},
	"ZuljinBuzzsaw": {
		Name: "Buzzsaw",
		Text: "After impact, Guillotine continues forward, dealing 75% damage to enemies hit. If a Hero is killed by Guillotine, Zul'jin is instantly healed to full.",
	},
	"ZuljinEnsnare": {
		Name: "Ensnare",
		Text: "Throw a net forward, Rooting the first enemy Hero hit for 1.5 seconds.",
	},
	"ZuljinEyeOfZuljin": {
		Name: "Eye of Zul'jin",
		Text: "Basic Attacks against enemy Heroes grant Zul'jin 6% Movement Speed for 2 seconds, up to 30%.",
	},
	"ZuljinFerocity": {
		Name: "Ferocity",
		Text: "Increases the Attack Speed bonus of Berserker by 40%.",
	},
	"ZuljinForestMedicine": {
		Name: "Forest Medicine",
		Text: "Regeneration is no longer Channeled, and cannot be interrupted.",
	},
	"ZuljinGuillotine": {
		Name: "Guillotine",
		Text: "Zul'jin launches a massive axe into the air that drops on the targeted area, dealing 350 damage plus bonus damage the lower his Health is.",
	},
	"ZuljinHeadhunter": {
		Name: "Headhunter",
		Text: "Quest: Zul'jin's first Takedown against each enemy Hero permanently increases his damage dealt by 2%.\nReward: After getting Takedowns on every enemy Hero, Basic Attack range is increased by 1.1.",
	},
	"ZuljinLacerate": {
		Name: "Lacerate",
		Text: "Increase the Slow amount of each Twin Cleave axe by 15%.",
	},
	"ZuljinLetTheKillingBegin": {
		Name: "Let the Killing Begin",
		Text: "When Zul'jin kills an enemy, his Attack Speed is increased by 5%, up to 40%, for 12 seconds. Basic Attacks refresh the duration.\nHero Takedowns instantly grant 40% Attack Speed.",
	},
	"ZuljinNoMercyTalent": {
		Name: "No Mercy!",
		Text: "Zul'jin's Basic Attacks against enemy Heroes marked with Grievous Throw ignore Armor.",
	},
	"ZuljinRecklessness": {
		Name: "Recklessness",
		Text: "While Zul'jin is below 75% Health, his Basic Attack damage is increased by 10%. While he is below 50% Health, he gains 20% Spell Power.",
	},
	"ZuljinSwirlingDeath": {
		Name: "Swirling Death",
		Text: "Twin Cleave's Axes cycle twice and can hit enemies on each revolution, but its cooldown is increased by 4 seconds.",
	},
	"ZuljinTazdingo": {
		Name: "Taz'dingo!",
		Text: "For the next 4 seconds, Zul'jin is Unkillable, and cannot be reduced to less than 1 Health. Taz'dingo!",
	},
	"ZuljinTrollsBlood": {
		Name: "Troll's Blood",
		Text: "Increase Regeneration's healing by 25%. If Regeneration fully finishes, restore 10% of Zul'jin's maximum Mana.",
	},
	"ZuljinViciousAssault": {
		Name: "Vicious Assault",
		Text: "Increase Grievous Throw's damage bonus by 35%.",
	},
	"ZuljinVoodooShuffle": {
		Name: "Voodoo Shuffle",
		Text: "Activate to remove Roots and Slows.\nPassive: Lower the cooldown and Mana cost of Regeneration by 33%.",
	},
	"ZuljinWrongPlaceWrongTime": {
		Name: "Wrong Place Wrong Time",
		Text: "If an enemy is hit by both Twin Cleave axes at the same time, they take an additional 150 damage and count as 5 Hero hits for You Want Axe?.",
	},
	"ZuljinYouWantAxe": {
		Name: "You Want Axe?",
		Text: "Quest: Every 5 Basic Attacks against Heroes permanently increases Basic Attack damage by 1.\nReward: After attacking Heroes 75 times, Basic Attack range is increased by 1.1.\nReward: After attacking Heroes 150 times, Twin Cleave now revolves twice.",
	},
}
