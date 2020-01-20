# NPCGEN - NON-Player Character Generator by Brandon Locker

PROJECT-0 (NPCGEN) is a CLI application which will assist a player in creating a non-player role-playing characters for use in table top RPGs. Currently following D20PFSRD rules found at https://www.d20pfsrd.com/classes/npc-classes/Non-player-characters/

## Help Menu
	This application is designed to help the user genrate basic NPCs.
	The core is designed around the D20 Pathfinder ruleset but some
	feature are available to allow for customization.
	Using the application with no argument will allow the user to specify the
	values based on answering prompts.
	
    -make     Generates a NPC using user defined variables.
	          variables are in the following order: sex name class
	          str dex con int wis cha
	          If values are left off they will be randomized.
	
              example: -make Male Fred Fighter Melee 13 11 12 9 10 8
	
    -random   Generates a fully random NPC using set defaults. If followed by a
	          value it will generate that many random NPCs
	
              example1: -random 5 || example2: -random
	
    -h, -help Displays this menu and exits the application\n")

    To provide more random results please add the following text files to the 
    same location as this application.

    bName.txt -> Text file holding boy names for randomization, one name per line.
    gName.txt -> Text file holding girl names for randomization, one name per line.
    class.txt -> Text file holding class names for randomization, one name per line.
    
## Future Plans
- [ ] Add way allow for creation and edit of text files
- [ ] Web / GUI interface
    