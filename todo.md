Outstanding
-----------
* Fix format so that it handles decimal places
* Verify the findbest algorithm is operating correctly
* Improve the findbest data structure
* Thread findbest to stop it from blocking the server.
* Saved message dissrupts moving from one field to another
	* need it to hover or something
* Refactor denormalizer event handlers
* Make repo a global variable? (command)
* Implement periodic snapshotting and event replay
* ?Only post if data actually changed?
	* I think using double-click to enter edit mode would work
* Add config files to servers
	* goinstall goconf.googlecode.com/hg
* Order the games coming out so they stay consistent
* Implement the scheduler
* Do some benchmarking

DONE
----
* Total does not update dynamically
* Have Finance and Structure accordians open by default
* Find out why the keys in the purchases are title case but not in the structures (query side)
* Get findbest results returning to browser
* merge monies and finance structs together
* Figure out why retain always is always false?!?!?
* Expanding purchases does not make the whole accordian div grow
* Make script to clean databases
* Create script to run all the servers
* Push code to bitbucket.org
* Implement event for changing structure cost
* Implement command for changing structure cost
* Some way to indicate that the cost field in the structure table are editable and the others are not
* Add structures
* Make denormalizer suspend on ctrl+z
	* make command and mongorestd do the same
	* CANNOT: limitation of Go signal handling
* Change rpc calls to use json instead of gob
* Make logger a global variable (command)
	* command/dispatcher.go not even using logger
* Error and success reporting on the page
	* Add "click to hide" text
	* Show success messages
* Add log files to servers
* Improve logging statements
* Use jquery toggleclass to show/hide shadow
* Get the tabs to autosize their height
* Figure out how to not have the nested divs animate when a section header is clicked
* Improve the layout and colours of the sections
* Add lands to game created commandhandler and eventhandler
* Change tasks to commands where appropriate
* Add the jQuery accordian to break up the page
* Use the jQuery tab ajax method to lazy-load the tabs
	* Refreshing data is refreshing all data, which is too broad
* Implement snapshot timestamping
* Destroying and recreating tabs causes first to be selected
* Add commandline flags
* Command snapshot bootstrap script
* Prevent enter from resizing the input boxes
* Refactor command handlers
* Refresh the data after a command is issued (always or should there be some kind of callback?)
* Snapshotting is not working, fix
* Refactor command handlers
* Fix the javascript currency formatting. numbers garbled again
* Re-add the game id to the posted command
* Implement event saving
* Implement jQuery tabs for the various games
* Reorganize website folder
* Refresh the data after a post
* Fix the denormalizer so it uses the game id for updates
* Add current date to task documents
* Pass the gameId to the command side
* Refactor index.js
* Figure out how to get numbers to transfer as numbers in JSON
* Validate numbers entered in javascript
* Have esc and enter change a field back to not having a shadow
* Have enter post data and make div not "focused"
* Have esc discard changes to a field
* Have lostfocus change a field back to not having a shadow
* Clean up the div entry fields so they don't get smaller when you delete or bigger when you fill them (they need to scroll the text if too big)
