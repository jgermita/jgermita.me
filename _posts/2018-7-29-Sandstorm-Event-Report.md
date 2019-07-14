!title: Sandstorm Event Report
!summary: My first combat robot event!

![](https://i.imgur.com/zwK8Iijl.jpg)

*Sandstorm Alpha(Left) and Beta(Right)*

Sandstorm performed fairly well for my first competitive entry in combat robotics. Destruction Under the Stars was an absolute blast and I'm quite happy with how things shook down. 

<iframe width="560" height="315" src="https://www.youtube.com/embed/NvqiugdCWNw?rel=0&amp;controls=0" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

*A short highlight reel of my favorite moments*

## Fight Summaries

### Sandstorm vs Gobble Gobble

<iframe width="560" height="315" src="https://www.youtube.com/embed/cfj5ilsl-es" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

The goal for this fight was to stay alive and land some good hits. Gobble Gobble was a standard beater bar bot that had the potential to take out my wheels or armor. The one thing I didn't expect going in was how small the arena was. I did not practice driving with any real space constraints and that was definitely an issue. 

The biggest issue with Sandstorm that ended up hurting in later fights was the lack of variable speed control on the weapon. Since I was spinning up to the full 9600+ RPM, I was storing so much KE that every turn I made had the potential to put me upside down. I spent a good part of the fight driving upside down until I landed a good hit that shot me into the arena ceiling and back on my feet. 

Other than those two issues, the fight went very well for Sandstorm. Win by KO.

### Sandstorm vs Impact

<iframe width="560" height="315" src="https://www.youtube.com/embed/lzIBo7Zs19Q?ecver=1" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

This was a scary fight going in. Impact is the defending champion with a very potent vertical drum spinner. In it's first fight of the day, it KO'd its opponent in under 10 seconds. With the corners I cut on armor design, I knew that any mistakes would be fatal. 

We traded some incredible weapon to weapon blows with Impact taking out one of my front armor plates early on. In the final moments of the fight, the weapon belt derailed causing me to lose my weapon. Luckily it derailed as it hit top speed, so I had about a second to make one last hit before it was too slow to do any damage. I was able to land a final hit that launched impact over the arena wall. Win by KO.

![](https://i.imgur.com/JjeBcH0l.jpg)

### Sandstorm vs Unknown Avenger

<iframe width="560" height="315" src="https://www.youtube.com/embed/x9qFKX6-FsM?ecver=1" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

Unknown Avenger is a well driven control bot. I knew that I had to keep my weapon functional for the entire fight to guarantee the W. Because of the weapon flipping issue, I knew that I couldn't win any judge decisions. Unknown Avenger capitalized on the time I spent getting back into fighting stance. He did not let up for the entire fight. As was the story with the last fight, I lost my weapon belt. I was unable to get into attacking position before it stopped and the fight turned into a wedge war. 

Additionally, I was fighting weird electrical issues all fight. My weapon was working just fine throughout, but I chose to spin it down at certain moments so I can reliably drive into position. Post mortem in the pit led me to believe that with the 1.5" wide wheels, I had too much traction so any time I hit a wall or robot, motors would stall vs wheels slipping. After an entire fight of stalling, they weren't performing quite as well towards the end. A quick glance after the rumble suggested that one of my LED strips was shorting out, potentially causing some other issues. I have yet to do a full diagnosis on the robot to identify any further issues.

Loss via Judge Decision.

### Sandstorm vs Duct Spartan

<iframe width="560" height="315" src="https://www.youtube.com/embed/UmIzGnn2G28?ecver=1" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

Duct Spartan is another aggressively driven control bot. The effective D2 drivetrain and wedges gave me lots of trouble throughout the entire match. For this match, I fielded the spare Sandstorm Beta with the thinner wheels and lack of LEDs. I was able to land some pretty good hits early on, but I lost my weapon for the third time. Duct Spartan simply out drove me. 

Loss via Judge Decision.

![](https://i.imgur.com/qY5nyRWl.jpg)

*Game Face*

### BONUS: Beetleweight Rumble

<iframe width="560" height="315" src="https://www.youtube.com/embed/a5s96Op8Jpw?ecver=1" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

This was the most absurdly fun robot thing I've ever done. 10 robots, 1 box, lots of fun. Sandstorm Alpha took to the box in this one. It was the last remaining KE weapon in the rumble mostly due to some cautious driving away from the pit door. I got greedy for kills and started driving closer and closer to the door when I got hung up by the antweight Meatball. We were both taken out by Burnt Toast Man and Smee.

![](https://i.imgur.com/wAW7m4Hl.jpg)

## TL;DR

### The Good

* Overall the system worked. Was able to drive pretty well and land some great hits.
* Never got KO'd 
* Zero fatal electrical issues
* Having a whole functional spare robot was really nice.

### The Bad

* Weapon belt reliability was a fatal issue. Ended up implementing a stupid fix before the rumble that I really should have identified weeks ago.
* Armor mounting was bad and I knew it. 
* Sandstorm stores too much KE for it's own good

### The Ugly

* Spent too much time on my side or upside down.
* 3D printed ABS frame failed exactly where expected. Luckily this didn't happen until the rumble. 

## Looking Forward

### Possible upgrades to Sandstorm

* Thinner or less massive weapon to free up weight. For the most part, it never used a huge chunk of the kinetic energy that it stored. 
* Variable speed control to weapon motor
* Better armor mounting
* 3D printed weapon belt retainer flange

### Future plans

* The two Sandstorm weapon bars will likely be repainted and used for Mirage, my horizontal spinner beetle design. 
* Antweights look fun. I want to build one. 
* If I can attend the 12lb event this fall, Echo will get a weapon system revamp to compete.

Thanks to everyone I met at Destruction Under the Stars. The combat robotics community is very welcoming and positive. Can't wait to see you all at the next one!

*Edit 7/31* More thoughts on what to change in designs moving forward

* Change batteries to XT30s. JST connectors that come stock on my batteries can only deliver about 10A continuous. XT30s can handle more, allowing me to use more of the battery's discharge current. 
* MAYBE make the battery smaller? Chargers showed I never really drained more than 25% of the battery's capacity per fight. This may change with the XT30 upgrade
* Top cover mounting screws to 6-32, use fewer screws. 4-40 heat set inserts were a little finicky to align, screws don't really come much longer than 1.5", and the clearance hole diameter is a little funky to cut on the CNC I use. 
* Put LEDs on a fused circuit just in case they get damaged. 
* Maybe do a small batch of press/screw on propdrive flanges for belt applications. 
* Do better testing vs wedges. Maybe build a wedge for testing? Maybe magnets would be a good idea?
* Single tooth weapon will probably be better for the Sandstorm design. 



