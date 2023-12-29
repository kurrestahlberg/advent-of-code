 Day 19 - Advent of Code 2023    window.addEventListener('click', function(e,s,r){if(e.target.nodeName==='CODE'&&e.detail===3){s=window.getSelection();s.removeAllRanges();r=document.createRange();r.selectNodeContents(e.target);s.addRange(r);}});

[Advent of Code](/)
===================

*   [\[About\]](/2023/about)
*   [\[Events\]](/2023/events)
*   [\[Shop\]](https://teespring.com/stores/advent-of-code)
*   [\[Settings\]](/2023/settings)
*   [\[Log Out\]](/2023/auth/logout)

Kurre Ståhlberg 36\*

  0.0.0.0:[2023](/2023)
=======================

*   [\[Calendar\]](/2023)
*   [\[AoC++\]](/2023/support)
*   [\[Sponsors\]](/2023/sponsors)
*   [\[Leaderboard\]](/2023/leaderboard)
*   [\[Stats\]](/2023/stats)

Our [sponsors](/2023/sponsors) help make Advent of Code possible:

[Optiver](https://optiver.com/advent-of-code) - Ready to solve puzzles for a living? We’re hiring C++, C# and Python experts to code sub-nanosecond trading systems. Get ready for problem solving, continuous learning and the freedom to bring your software solutions to life

\--- Day 19: Aplenty ---
------------------------

The Elves of Gear Island are thankful for your help and send you on your way. They even have a hang glider that someone [stole](9) from Desert Island; since you're already going that direction, it would help them a lot if you would use it to get down there and return it to them.

As you reach the bottom of the _relentless avalanche of machine parts_, you discover that they're already forming a formidable heap. Don't worry, though - a group of Elves is already here organizing the parts, and they have a _system_.

To start, each part is rated in each of four categories:

*   `x`: E_x_tremely cool looking
*   `m`: _M_usical (it makes a noise when you hit it)
*   `a`: _A_erodynamic
*   `s`: _S_hiny

Then, each part is sent through a series of _workflows_ that will ultimately _accept_ or _reject_ the part. Each workflow has a name and contains a list of _rules_; each rule specifies a condition and where to send the part if the condition is true. The first rule that matches the part being considered is applied immediately, and the part moves on to the destination described by the rule. (The last rule in each workflow has no condition and always applies if reached.)

Consider the workflow `ex{x>10:one,m<20:two,a>30:R,A}`. This workflow is named `ex` and contains four rules. If workflow `ex` were considering a specific part, it would perform the following steps in order:

*   Rule "`x>10:one`": If the part's `x` is more than `10`, send the part to the workflow named `one`.
*   Rule "`m<20:two`": Otherwise, if the part's `m` is less than `20`, send the part to the workflow named `two`.
*   Rule "`a>30:R`": Otherwise, if the part's `a` is more than `30`, the part is immediately _rejected_ (`R`).
*   Rule "`A`": Otherwise, because no other rules matched the part, the part is immediately _accepted_ (`A`).

If a part is sent to another workflow, it immediately switches to the start of that workflow instead and never returns. If a part is _accepted_ (sent to `A`) or _rejected_ (sent to `R`), the part immediately stops any further processing.

The system works, but it's not keeping up with the torrent of weird metal shapes. The Elves ask if you can help sort a few parts and give you the list of workflows and some part ratings (your puzzle input). For example:

    px{a<2006:qkq,m>2090:A,rfg}
    pv{a>1716:R,A}
    lnx{m>1548:A,A}
    rfg{s<537:gd,x>2440:R,A}
    qs{s>3448:A,lnx}
    qkq{x<1416:A,crn}
    crn{x>2662:A,R}
    in{s<1351:px,qqz}
    qqz{s>2770:qs,m<1801:hdj,R}
    gd{a>3333:R,R}
    hdj{m>838:A,pv}
    
    {x=787,m=2655,a=1222,s=2876}
    {x=1679,m=44,a=2067,s=496}
    {x=2036,m=264,a=79,s=2244}
    {x=2461,m=1339,a=466,s=291}
    {x=2127,m=1623,a=2188,s=1013}
    

The workflows are listed first, followed by a blank line, then the ratings of the parts the Elves would like you to sort. All parts begin in the workflow named `in`. In this example, the five listed parts go through the following workflows:

*   `{x=787,m=2655,a=1222,s=2876}`: `in` -> `qqz` -> `qs` -> `lnx` -> `_A_`
*   `{x=1679,m=44,a=2067,s=496}`: `in` -> `px` -> `rfg` -> `gd` -> `_R_`
*   `{x=2036,m=264,a=79,s=2244}`: `in` -> `qqz` -> `hdj` -> `pv` -> `_A_`
*   `{x=2461,m=1339,a=466,s=291}`: `in` -> `px` -> `qkq` -> `crn` -> `_R_`
*   `{x=2127,m=1623,a=2188,s=1013}`: `in` -> `px` -> `rfg` -> `_A_`

Ultimately, three parts are _accepted_. Adding up the `x`, `m`, `a`, and `s` rating for each of the accepted parts gives `7540` for the part with `x=787`, `4623` for the part with `x=2036`, and `6951` for the part with `x=2127`. Adding all of the ratings for _all_ of the accepted parts gives the sum total of `_19114_`.

Sort through all of the parts you've been given; _what do you get if you add together all of the rating numbers for all of the parts that ultimately get accepted?_

To begin, [get your puzzle input](19/input).

Answer:  

You can also \[Shareon [Twitter](https://twitter.com/intent/tweet?text=%22Aplenty%22+%2D+Day+19+%2D+Advent+of+Code+2023&url=https%3A%2F%2Fadventofcode%2Ecom%2F2023%2Fday%2F19&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)\] this puzzle.

(function(i,s,o,g,r,a,m){i\['GoogleAnalyticsObject'\]=r;i\[r\]=i\[r\]||function(){ (i\[r\].q=i\[r\].q||\[\]).push(arguments)},i\[r\].l=1\*new Date();a=s.createElement(o), m=s.getElementsByTagName(o)\[0\];a.async=1;a.src=g;m.parentNode.insertBefore(a,m) })(window,document,'script','//www.google-analytics.com/analytics.js','ga'); ga('create', 'UA-69522494-1', 'auto'); ga('set', 'anonymizeIp', true); ga('send', 'pageview');