 Day 18 - Advent of Code 2023    window.addEventListener('click', function(e,s,r){if(e.target.nodeName==='CODE'&&e.detail===3){s=window.getSelection();s.removeAllRanges();r=document.createRange();r.selectNodeContents(e.target);s.addRange(r);}});

[Advent of Code](/)
===================

*   [\[About\]](/2023/about)
*   [\[Events\]](/2023/events)
*   [\[Shop\]](https://teespring.com/stores/advent-of-code)
*   [\[Settings\]](/2023/settings)
*   [\[Log Out\]](/2023/auth/logout)

Kurre Ståhlberg 34\*

      /\*[2023](/2023)\*/
=========================

*   [\[Calendar\]](/2023)
*   [\[AoC++\]](/2023/support)
*   [\[Sponsors\]](/2023/sponsors)
*   [\[Leaderboard\]](/2023/leaderboard)
*   [\[Stats\]](/2023/stats)

Our [sponsors](/2023/sponsors) help make Advent of Code possible:

[Cerbos](https://bit.ly/3MPonyQ) - Easily implement and manage fine-grained access control in your app

\--- Day 18: Lavaduct Lagoon ---
--------------------------------

Thanks to your efforts, the machine parts factory is one of the first factories up and running since the lavafall came back. However, to catch up with the large backlog of parts requests, the factory will also need a _large supply of lava_ for a while; the Elves have already started creating a large lagoon nearby for this purpose.

However, they aren't sure the lagoon will be big enough; they've asked you to take a look at the _dig plan_ (your puzzle input). For example:

    R 6 (#70c710)
    D 5 (#0dc571)
    L 2 (#5713f0)
    D 2 (#d2c081)
    R 2 (#59c680)
    D 2 (#411b91)
    L 5 (#8ceee2)
    U 2 (#caa173)
    L 1 (#1b58a2)
    U 2 (#caa171)
    R 2 (#7807d2)
    U 3 (#a77fa3)
    L 2 (#015232)
    U 2 (#7a21e3)
    

The digger starts in a 1 meter cube hole in the ground. They then dig the specified number of meters _up_ (`U`), _down_ (`D`), _left_ (`L`), or _right_ (`R`), clearing full 1 meter cubes as they go. The directions are given as seen from above, so if "up" were north, then "right" would be east, and so on. Each trench is also listed with _the color that the edge of the trench should be painted_ as an [RGB hexadecimal color code](https://en.wikipedia.org/wiki/RGB_color_model#Numeric_representations).

When viewed from above, the above example dig plan would result in the following loop of _trench_ (`#`) having been dug out from otherwise _ground-level terrain_ (`.`):

    #######
    #.....#
    ###...#
    ..#...#
    ..#...#
    ###.###
    #...#..
    ##..###
    .#....#
    .######
    

At this point, the trench could contain 38 cubic meters of lava. However, this is just the edge of the lagoon; the next step is to _dig out the interior_ so that it is one meter deep as well:

    #######
    #######
    #######
    ..#####
    ..#####
    #######
    #####..
    #######
    .######
    .######
    

Now, the lagoon can contain a much more respectable `_62_` cubic meters of lava. While the interior is dug out, the edges are also painted according to the color codes in the dig plan.

The Elves are concerned the lagoon won't be large enough; if they follow their dig plan, _how many cubic meters of lava could it hold?_

To begin, [get your puzzle input](18/input).

Answer:  

You can also \[Shareon [Twitter](https://twitter.com/intent/tweet?text=%22Lavaduct+Lagoon%22+%2D+Day+18+%2D+Advent+of+Code+2023&url=https%3A%2F%2Fadventofcode%2Ecom%2F2023%2Fday%2F18&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)\] this puzzle.

(function(i,s,o,g,r,a,m){i\['GoogleAnalyticsObject'\]=r;i\[r\]=i\[r\]||function(){ (i\[r\].q=i\[r\].q||\[\]).push(arguments)},i\[r\].l=1\*new Date();a=s.createElement(o), m=s.getElementsByTagName(o)\[0\];a.async=1;a.src=g;m.parentNode.insertBefore(a,m) })(window,document,'script','//www.google-analytics.com/analytics.js','ga'); ga('create', 'UA-69522494-1', 'auto'); ga('set', 'anonymizeIp', true); ga('send', 'pageview');