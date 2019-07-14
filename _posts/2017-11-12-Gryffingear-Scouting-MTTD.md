!title: Gryffingear Scouting at the 2017 Madtown Throwdown
!summary: A breakdown of Gryffingear's Google Sheets based scouting system

This post will look at the key features of Gryffingear's scouting system for the 2017 Madtown Throwdown. This scouting system was integral in making informed strategy decisions throughout qualification matches, during alliance selections, and into eliminations matches. 

The scouting system is built on the groundwork laid by Gryffingear alumni Brian Romero, Javier Serna, and Christian Perry during the 2017 competition season thanks to the assistance and advice from numerous other teams. 

# Goals

The goals of this system were as follows:

1. Collects only the data that matters - quantify qualities that most teams should be more or less proficient at (# of gears scored in matches in 2017), collect yes/no's for boolean qualities (climbing) or "icing on the cake" qualities that fewer teams have (ability to hit 40KPA during the match).
2. Make data entry dead simple - Don't make it a chore to get data from the scouts to the scouting system.
3. Automate as much analysis as possible - matches are played far faster than any human can compile data and produce meaningful metrics and analysis. Automate this data analysis and aggregation.

Goal #1 was accomplished through analysis of how we've played the game thus far. After playing through this game at several events, we have a good grasp at what data is useful for strategic decisions and what is just noise. In the future, this would be accomplished through effective strategic analysis of the game. 

Goal #2 was simple. 2017 competition season saw the rise of Google Forms in our scouting workflow. Google Forms allows us to not reinvent the data-entry wheel by riding on Google's massive infrastructure to collect our data from multiple sources (scouts). The one caveat to our solution is that it relies on scouts using their own data connections to do their jobs. Luckily Forms doesn't transfer too much data up or down, but the problem isn't totally eliminated.

Goal #3 was accomplished using Google Sheets and the powerful formula and query language baked in. Thanks to assistance and advice from various teams, our competition season scouting system further bridged the gap between the data entered and the data aggregated and analyzed.

The MTTD iteration of this system builds on our competition season's solutions to these problems. 

# Overview

## Data Entry

![Google Form](https://i.imgur.com/K3ftppqm.png)

Google Forms is used heavily in data entry. The link to the scouting form for a specific event is shared with the scouting team at the beginning of each event. 

![Raw data in Google Sheet](https://i.imgur.com/pgAJbWcl.png)

After entry, data is automatically sent by Forms to a Google Sheet. All further data analysis references this one sheet. 

![Pit scouting data](https://i.imgur.com/b3o71zql.png)

If needed, pit scouting data like drivetrain type, mechanisms, and robot pictures are also entered into a separate sheet in the same Google Sheets document.

## Analysis and Aggregation

![all team data](https://i.imgur.com/qZjtY2Il.png)

Data on each specific team is eventually summarized by the AllTeamData sheet. This takes all scouted match data and generates metrics such as min, max, and averages on all values. All magic happens thanks to this sheet. 

![team summary sheet](https://i.imgur.com/ZXy4imPl.png)

This view is what I used most frequently as a coach. It uses data from the AllTeamData and Pit Scouting sheets and presents it in large and easy to read format (semi-optimized for mobile) so that I can run queries on specific teams and quickly bring up their data. What used to take hours under our 2016 scouting system took a mere 5 seconds with the 2017 offseason system.

![auto picklist](https://i.imgur.com/YwmLyApl.png)

The next most useful feature of the new system. By using the ATD sheet, each team is scored on a weighted sum of their average gears, climb percentage, died percentage, and other metrics. Doing this, we can automatically generate a pick list to use during alliance selections. Our third pick ended up being a team from the APL. The APL view also has cells that allow us to indicate whether or not teams are picked or should be skipped so that the next major component of the system can work.

![Whiteboard View](https://i.imgur.com/Xr3VgxAl.png)

Whiteboard View! Filling a long-time need in FRC to have a clear, easy to read, "whiteboard" type view, this allows us to communicate to our alliance selections representative from the stands. This was brought up on a 10.1" tablet and displayed to our reps from over 50ft away. This view also has options for "ACCEPT", "DENY", or "DEFER" so that we can fully communicate to our representative. Inspiration for this was drawn by numerous mishaps by my teams as well as countless others to clearly communicate these on-the-fly picks to their rep.

A short demo of APL and Whiteboard are shown below:

<iframe width="560" height="315" src="https://www.youtube.com/embed/iK8YCYPBgzU?rel=0&amp;showinfo=0" frameborder="0" allowfullscreen></iframe>

# Under the Hood

The Google Sheet heavily relies on Google's implementation of [queries](https://support.google.com/docs/answer/3093343?hl=en) to do all of the heavy lifting. The language used by Google Sheets is a stripped down version of common database query languages like SQL. Using common database queries, we can quickly reference data dynamically wherever we'd need. Using a combination of queries and the powerful equation engine of Google Sheets, we were able to automate 100% of the data aggregation and analysis done. 

Here's the equation that forms the backbone of the Team Summary sheet. It queries the form response data to show basic stats on a team's performance match by match.

`=IFERROR(IF(len(D1)=0,"Please choose a value in D1",query(team_data!A5:S31,"Select C, M, N, O, P, Q, R, S, L Where D = "&D1&" order by C")),)`

The `IFERROR()` is a nice way to keep the cells blank if no data is found. Otherwise, they'd be populated with zeros or other error text. `IF(len(D10)=0` checks to see if the input cell(the team to query for) is empty and produces an error message if it is empty. 

The real meat of this equation is in the `query()`, stripped down below:

`query(team_data!A5:S31,"Select C, M, N, O, P, Q, R, S, L Where D = "&D1&" order by C")`

The first argument is the dataset to query and the second argument is the query itself. It will return columns C, M, N, O, P, Q, R, S, and L where column D in the dataset is equal to cell D1, then sorts this new dataset by column C (match number).

Here's the output:

![ ](https://i.imgur.com/PFkM0Fal.png)

From here, we now have a nice looking set of data to perform `average()`, `min()`, `max()`, or other equations on:

![](https://i.imgur.com/TRDF9Oyl.png)

# Improvements

This system was a massive step forward for Gryffingear scouting, but it is far from perfect. During the event, we encountered multiple bugs. These would have been mitigated by more intensive pre-event testing. 
There were minor inconsistencies in how the users entered data, but that could be mitigated in the future by proper training and testing on the system.
Potentially, this sheet could just expose the data as a database so that we could develop a more focused application such as Tableau or some other custom app to view the data. 

Moving forward, an elims view would also be good as would be a way to view elims matches in the match summary view.

You can view the entire system as of the end of the event [here.](https://docs.google.com/spreadsheets/d/1sgMFYy1W4-KNFZGhsyPzQ1Qby11oNo8Zl0JMXrbK8z8/edit?usp=sharing)

# Acknowledgements

This system was inspired by efforts and techniques by teams 987, 973, 2659, 1678, and numerous other inspiring teams we've been fortunate enough to compete with and against over the years. Scouting can make or break a competition for teams and these teams show year in and out how to fully leverage the data to consistently win matches and events. 

Gryffingear class of 2017 alumni Brian Romero, Javier Serna, and Christian Perry laid solid groundwork for much of this system during the competition season. The effects of their efforts will be felt for years to come. 

A tool is only as useful as its users. Our scouting team at Madtown Throwdown 2017 did a fantastic job at entering quality data on each and every match. Scouting is not a glamorous job, nor is it always fun. I'm grateful to them for putting in the effort to support the team's strategic efforts on and off the field. 

# Further Reading

An understanding of Google sheets equations and queries are critical in buildinga scouting system of this scale. Below are a couple of links that dive into topics relevant to the techniques used in this scouting system. 

* [Equations tutorial](https://www.thoughtco.com/google-spreadsheet-formula-tutorial-3123950)
* [Tutorial on Query](https://codingisforlosers.com/google-sheets-query-function/)
* [Index-match](http://www.mbaexcel.com/excel/why-index-match-is-better-than-vlookup/) - Faster than vlookup

