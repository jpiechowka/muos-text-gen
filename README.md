# muOS Text Gen

A little utility to generate text descriptions for muOS for games scraped with Skraper.

* https://muos.dev/
* https://muos.dev/help/addcontent

## Usage
1. Scrape your content with Skraper (https://www.skraper.net/)
2. Copy the binary to the directory that contains your generated `.dat` file (or `.xml` if renamed)
3. Run the binary (it will create `text` directory in the folder the binary was run)
4. Follow the guide here to copy your text descriptions https://muos.dev/help/addcontent

## Sample run
```
2024/06/12 16:14:12.788672 main.go:51: Current operation directory: D:\Roms\Nintendo - Game Boy Advance
2024/06/12 16:14:12.789812 main.go:61: Reading file: D:\Roms\Nintendo - Game Boy Advance\Game Boy Advance.dat
2024/06/12 16:14:12.791446 main.go:103: Generating description for: Kirby : Nightmare in Dream Land
2024/06/12 16:14:12.791985 main.go:123: Generated description for Kirby : Nightmare in Dream Land in: text/Kirby - Nightmare in Dream Land (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.791985 main.go:103: Generating description for: Mario & Luigi : Superstar Saga
2024/06/12 16:14:12.791985 main.go:123: Generated description for Mario & Luigi : Superstar Saga in: text/Mario & Luigi - Superstar Saga (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.791985 main.go:103: Generating description for: Super Mario World: Super Mario Advance 2
2024/06/12 16:14:12.791985 main.go:123: Generated description for Super Mario World: Super Mario Advance 2 in: text/Super Mario Advance 2 - Super Mario World (Europe) (En,Fr,De,Es).gba.txt
2024/06/12 16:14:12.791985 main.go:103: Generating description for: Advance Wars
2024/06/12 16:14:12.792502 main.go:123: Generated description for Advance Wars in: text/Advance Wars (Europe) (En,Fr,De,Es).gba.txt
2024/06/12 16:14:12.792502 main.go:103: Generating description for: Metroid Fusion
2024/06/12 16:14:12.792502 main.go:123: Generated description for Metroid Fusion in: text/Metroid Fusion (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.792502 main.go:103: Generating description for: Pokémon Emerald Version
2024/06/12 16:14:12.792502 main.go:123: Generated description for Pokémon Emerald Version in: text/Pokemon - Emerald Version (USA, Europe).gba.txt
2024/06/12 16:14:12.792502 main.go:103: Generating description for: Metroid : Zero Mission
2024/06/12 16:14:12.793011 main.go:123: Generated description for Metroid : Zero Mission in: text/Metroid - Zero Mission (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.793011 main.go:103: Generating description for: Dragon Ball : Advanced Adventure
2024/06/12 16:14:12.793011 main.go:123: Generated description for Dragon Ball : Advanced Adventure in: text/Dragon Ball - Advanced Adventure (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.793011 main.go:103: Generating description for: Sonic Pinball Party
2024/06/12 16:14:12.793532 main.go:123: Generated description for Sonic Pinball Party in: text/Sonic Pinball Party (Europe) (En,Ja,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.793532 main.go:103: Generating description for: The Legend of Zelda : A Link to the Past & Four Sw
2024/06/12 16:14:12.793532 main.go:123: Generated description for The Legend of Zelda : A Link to the Past & Four Sw in: text/Legend of Zelda, The - A Link to the Past & Four Swords (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.793532 main.go:103: Generating description for: Crash Bandicoot Fusion
2024/06/12 16:14:12.794050 main.go:123: Generated description for Crash Bandicoot Fusion in: text/Crash Bandicoot Fusion (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.794050 main.go:103: Generating description for: Wario Land 4
2024/06/12 16:14:12.794050 main.go:123: Generated description for Wario Land 4 in: text/Wario Land 4 (USA, Europe).gba.txt
2024/06/12 16:14:12.794050 main.go:103: Generating description for: TMNT : Teenage Mutant Ninja Turtles
2024/06/12 16:14:12.794050 main.go:123: Generated description for TMNT : Teenage Mutant Ninja Turtles in: text/TMNT - Teenage Mutant Ninja Turtles (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.794050 main.go:103: Generating description for: Crash Nitro Kart
2024/06/12 16:14:12.794589 main.go:123: Generated description for Crash Nitro Kart in: text/Crash Nitro Kart (Europe) (En,Fr,De,Es,It,Nl).gba.txt
2024/06/12 16:14:12.794589 main.go:103: Generating description for: F-Zero : GP Legend
2024/06/12 16:14:12.794589 main.go:123: Generated description for F-Zero : GP Legend in: text/F-Zero - GP Legend (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.794589 main.go:103: Generating description for: Castlevania
2024/06/12 16:14:12.794589 main.go:123: Generated description for Castlevania in: text/Castlevania (Europe).gba.txt
2024/06/12 16:14:12.795113 main.go:103: Generating description for: Mario Kart : Super Circuit
2024/06/12 16:14:12.795113 main.go:123: Generated description for Mario Kart : Super Circuit in: text/Mario Kart - Super Circuit (Europe).gba.txt
2024/06/12 16:14:12.795113 main.go:103: Generating description for: Castlevania : Aria of Sorrow
2024/06/12 16:14:12.795113 main.go:123: Generated description for Castlevania : Aria of Sorrow in: text/Castlevania - Aria of Sorrow (Europe) (En,Fr,De).gba.txt
2024/06/12 16:14:12.795113 main.go:103: Generating description for: Advance Wars 2 : Black Hole Rising
2024/06/12 16:14:12.795113 main.go:123: Generated description for Advance Wars 2 : Black Hole Rising in: text/Advance Wars 2 - Black Hole Rising (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.795113 main.go:103: Generating description for: Sonic Battle
2024/06/12 16:14:12.795632 main.go:123: Generated description for Sonic Battle in: text/Sonic Battle (Europe) (En,Ja,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.795632 main.go:103: Generating description for: Max Payne Advance
2024/06/12 16:14:12.795632 main.go:123: Generated description for Max Payne Advance in: text/Max Payne Advance (Europe) (En,Fr,De).gba.txt
2024/06/12 16:14:12.795632 main.go:103: Generating description for: Crash Bandicoot 2 : N-Tranced
2024/06/12 16:14:12.795632 main.go:123: Generated description for Crash Bandicoot 2 : N-Tranced in: text/Crash Bandicoot 2 - N-Tranced (Europe) (En,Fr,De,Es,It,Nl).gba.txt
2024/06/12 16:14:12.795632 main.go:103: Generating description for: F-Zero : Maximum Velocity
2024/06/12 16:14:12.796151 main.go:123: Generated description for F-Zero : Maximum Velocity in: text/F-Zero - Maximum Velocity (USA, Europe).gba.txt
2024/06/12 16:14:12.796151 main.go:103: Generating description for: Medal of Honor : Infiltrator
2024/06/12 16:14:12.796151 main.go:123: Generated description for Medal of Honor : Infiltrator in: text/Medal of Honor - Infiltrator (USA, Europe) (En,Fr,De).gba.txt
2024/06/12 16:14:12.796151 main.go:103: Generating description for: V-Rally 3
2024/06/12 16:14:12.796151 main.go:123: Generated description for V-Rally 3 in: text/V-Rally 3 (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.796151 main.go:103: Generating description for: Metal Slug Advance
2024/06/12 16:14:12.796151 main.go:123: Generated description for Metal Slug Advance in: text/Metal Slug Advance (Europe).gba.txt
2024/06/12 16:14:12.796151 main.go:103: Generating description for: Final Fantasy VI Advance
2024/06/12 16:14:12.796151 main.go:123: Generated description for Final Fantasy VI Advance in: text/Final Fantasy VI Advance (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.796151 main.go:103: Generating description for: Sonic Advance 3
2024/06/12 16:14:12.796668 main.go:123: Generated description for Sonic Advance 3 in: text/Sonic Advance 3 (Europe) (En,Ja,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.796668 main.go:103: Generating description for: The Legend of Zelda : The Minish Cap
2024/06/12 16:14:12.796668 main.go:123: Generated description for The Legend of Zelda : The Minish Cap in: text/Legend of Zelda, The - The Minish Cap (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.796668 main.go:103: Generating description for: Mario Party Advance
2024/06/12 16:14:12.796668 main.go:123: Generated description for Mario Party Advance in: text/Mario Party Advance (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.796668 main.go:103: Generating description for: Mega Man Zero 3
2024/06/12 16:14:12.796668 main.go:123: Generated description for Mega Man Zero 3 in: text/Megaman Zero 3 (Europe).gba.txt
2024/06/12 16:14:12.796668 main.go:103: Generating description for: Yoshi's Island: Super Mario Advance 3
2024/06/12 16:14:12.796668 main.go:123: Generated description for Yoshi's Island: Super Mario Advance 3 in: text/Super Mario Advance 3 - Yoshi's Island (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.797206 main.go:103: Generating description for: Sonic Advance 2
2024/06/12 16:14:12.797206 main.go:123: Generated description for Sonic Advance 2 in: text/Sonic Advance 2 (Europe) (En,Ja,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.797206 main.go:103: Generating description for: Iridion II
2024/06/12 16:14:12.797206 main.go:123: Generated description for Iridion II in: text/Iridion II (Europe) (En,Fr,De).gba.txt
2024/06/12 16:14:12.797206 main.go:103: Generating description for: WarioWare, Inc. : Minigame Mania
2024/06/12 16:14:12.797206 main.go:123: Generated description for WarioWare, Inc. : Minigame Mania in: text/WarioWare, Inc. - Minigame Mania (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.797206 main.go:103: Generating description for: Final Fantasy Tactics Advance
2024/06/12 16:14:12.797206 main.go:123: Generated description for Final Fantasy Tactics Advance in: text/Final Fantasy Tactics Advance (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.797741 main.go:103: Generating description for: Virtua Tennis
2024/06/12 16:14:12.797741 main.go:123: Generated description for Virtua Tennis in: text/Virtua Tennis (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.797741 main.go:103: Generating description for: Crash Bandicoot XS
2024/06/12 16:14:12.797741 main.go:123: Generated description for Crash Bandicoot XS in: text/Crash Bandicoot XS (Europe) (En,Fr,De,Es,It,Nl).gba.txt
2024/06/12 16:14:12.797741 main.go:103: Generating description for: Kirby & the Amazing Mirror
2024/06/12 16:14:12.797741 main.go:123: Generated description for Kirby & the Amazing Mirror in: text/Kirby & The Amazing Mirror (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.797741 main.go:103: Generating description for: Ninja Cop
2024/06/12 16:14:12.798255 main.go:123: Generated description for Ninja Cop in: text/Ninja Cop (Europe).gba.txt
2024/06/12 16:14:12.798255 main.go:103: Generating description for: Mario Power Tennis
2024/06/12 16:14:12.798255 main.go:123: Generated description for Mario Power Tennis in: text/Mario Power Tennis (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.798255 main.go:103: Generating description for: Mega Man Zero
2024/06/12 16:14:12.798255 main.go:123: Generated description for Mega Man Zero in: text/Megaman Zero (USA, Europe).gba.txt
2024/06/12 16:14:12.798255 main.go:103: Generating description for: Colin McRae Rally 2.0
2024/06/12 16:14:12.798255 main.go:123: Generated description for Colin McRae Rally 2.0 in: text/Colin McRae Rally 2.0 (Europe) (En,Fr,De).gba.txt
2024/06/12 16:14:12.798255 main.go:103: Generating description for: Drill Dozer
2024/06/12 16:14:12.798255 main.go:123: Generated description for Drill Dozer in: text/Drill Dozer (USA).gba.txt
2024/06/12 16:14:12.798255 main.go:103: Generating description for: Super Mario Advance 4: Super Mario Bros. 3
2024/06/12 16:14:12.798773 main.go:123: Generated description for Super Mario Advance 4: Super Mario Bros. 3 in: text/Super Mario Advance 4 - Super Mario Bros. 3 (Europe) (En,Fr,De,Es,It) (Rev 1).gba.txt
2024/06/12 16:14:12.798773 main.go:103: Generating description for: Super Mario Ball
2024/06/12 16:14:12.798773 main.go:123: Generated description for Super Mario Ball in: text/Super Mario Ball (Europe).gba.txt
2024/06/12 16:14:12.798773 main.go:103: Generating description for: Street Fighter Alpha 3
2024/06/12 16:14:12.798773 main.go:123: Generated description for Street Fighter Alpha 3 in: text/Street Fighter Alpha 3 (Europe).gba.txt
2024/06/12 16:14:12.798773 main.go:103: Generating description for: Mega Man Zero 4
2024/06/12 16:14:12.798773 main.go:123: Generated description for Mega Man Zero 4 in: text/Megaman Zero 4 (Europe).gba.txt
2024/06/12 16:14:12.798773 main.go:103: Generating description for: Super Mario Advance
2024/06/12 16:14:12.799290 main.go:123: Generated description for Super Mario Advance in: text/Super Mario Advance (USA, Europe).gba.txt
2024/06/12 16:14:12.799290 main.go:103: Generating description for: Mega Man Zero 2
2024/06/12 16:14:12.799290 main.go:123: Generated description for Mega Man Zero 2 in: text/Megaman Zero 2 (Europe).gba.txt
2024/06/12 16:14:12.799290 main.go:103: Generating description for: Mario vs. Donkey Kong
2024/06/12 16:14:12.799290 main.go:123: Generated description for Mario vs. Donkey Kong in: text/Mario vs. Donkey Kong (Europe) (En,Fr,De,Es,It).gba.txt
2024/06/12 16:14:12.799290 main.go:103: Generating description for: Sonic Advance
2024/06/12 16:14:12.799808 main.go:123: Generated description for Sonic Advance in: text/Sonic Advance (Europe) (En,Ja,Fr,De,Es).gba.txt
2024/06/12 16:14:12.799808 main.go:103: Generating description for: Castlevania - Harmony Of Dissonance
2024/06/12 16:14:12.799808 main.go:123: Generated description for Castlevania - Harmony Of Dissonance in: text/Castlevania - Harmony of Dissonance (Europe).gba.txt
2024/06/12 16:14:12.799808 main.go:103: Generating description for: NES Classics : Super Mario Bros.
2024/06/12 16:14:12.799808 main.go:123: Generated description for NES Classics : Super Mario Bros. in: text/Classic NES Series - Super Mario Bros. (USA, Europe).gba.txt
2024/06/12 16:14:12.799808 main.go:103: Generating description for: Mario Golf : Advance Tour
2024/06/12 16:14:12.800328 main.go:123: Generated description for Mario Golf : Advance Tour in: text/Mario Golf - Advance Tour (Europe).gba.txt
2024/06/12 16:14:12.800328 main.go:103: Generating description for: Pokémon FireRed Version
2024/06/12 16:14:12.800328 main.go:123: Generated description for Pokémon FireRed Version in: text/Pokemon - FireRed Version (USA, Europe) (Rev 1).gba.txt
2024/06/12 16:14:12.800328 main.go:126: Finished generating description for 56 games
```