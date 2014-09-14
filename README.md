# Best Buys #

Overview
--------
Best Buys is an application for determining which lands/bases to buy, and in what order, for Uken's [rpg games](http://uken.com/#games). Given the current purchase price of your lands, your hourly income and your bank account balance, it will run a simluation of purchasing your next six lands. Using the results of the simulation, it will find the order of purchases that will maximize your income at the end of the six purchases.

Technology
----------
- The front-end UI is an HTML5 and javascript single page appliation. It using client-side templating and ajax.
- The back-end uses Go (Google's golang) and MongoDB to produce a REST API that the front-end uses.
- The architecture of the system is CQRS (command query responsibility segregation). Event sourcing is employed such that all changes made in the UI are captured in an event database.