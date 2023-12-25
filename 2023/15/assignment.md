 Day 15 - Advent of Code 2023    window.addEventListener('click', function(e,s,r){if(e.target.nodeName==='CODE'&&e.detail===3){s=window.getSelection();s.removeAllRanges();r=document.createRange();r.selectNodeContents(e.target);s.addRange(r);}});

[Advent of Code](/)
===================

*   [\[About\]](/2023/about)
*   [\[Events\]](/2023/events)
*   [\[Shop\]](https://teespring.com/stores/advent-of-code)
*   [\[Settings\]](/2023/settings)
*   [\[Log Out\]](/2023/auth/logout)

Kurre Ståhlberg 28\*

       y([2023](/2023))
=======================

*   [\[Calendar\]](/2023)
*   [\[AoC++\]](/2023/support)
*   [\[Sponsors\]](/2023/sponsors)
*   [\[Leaderboard\]](/2023/leaderboard)
*   [\[Stats\]](/2023/stats)

Our [sponsors](/2023/sponsors) help make Advent of Code possible:

[Splunk](https://www.splunk.com/en_us/careers.html) - Come build a more resilient digital world with us.

\--- Day 15: Lens Library ---
-----------------------------

The newly-focused parabolic reflector dish is sending all of the collected light to a point on the side of yet another mountain - the largest mountain on Lava Island. As you approach the mountain, you find that the light is being collected by the wall of a large facility embedded in the mountainside.

You find a door under a large sign that says "Lava Production Facility" and next to a smaller sign that says "Danger - Personal Protective Equipment required beyond this point".

As you step inside, you are immediately greeted by a somewhat panicked reindeer wearing goggles and a loose-fitting [hard hat](https://en.wikipedia.org/wiki/Hard_hat). The reindeer leads you to a shelf of goggles and hard hats (you quickly find some that fit) and then further into the facility. At one point, you pass a button with a faint snout mark and the label "PUSH FOR HELP". No wonder you were loaded into that [trebuchet](1) so quickly!

You pass through a final set of doors surrounded with even more warning signs and into what must be the room that collects all of the light from outside. As you admire the large assortment of lenses available to further focus the light, the reindeer brings you a book titled "Initialization Manual".

"Hello!", the book cheerfully begins, apparently unaware of the concerned reindeer reading over your shoulder. "This procedure will let you bring the Lava Production Facility online - all without burning or melting anything unintended!"

"Before you begin, please be prepared to use the Holiday ASCII String Helper algorithm (appendix 1A)." You turn to appendix 1A. The reindeer leans closer with interest.

The HASH algorithm is a way to turn any [string](https://en.wikipedia.org/wiki/String_(computer_science)) of characters into a single _number_ in the range 0 to 255. To run the HASH algorithm on a string, start with a _current value_ of `0`. Then, for each character in the string starting from the beginning:

*   Determine the [ASCII code](https://en.wikipedia.org/wiki/ASCII#Printable_characters) for the current character of the string.
*   Increase the _current value_ by the ASCII code you just determined.
*   Set the _current value_ to itself multiplied by `17`.
*   Set the _current value_ to the [remainder](https://en.wikipedia.org/wiki/Modulo) of dividing itself by `256`.

After following these steps for each character in the string in order, the _current value_ is the output of the HASH algorithm.

So, to find the result of running the HASH algorithm on the string `HASH`:

*   The _current value_ starts at `0`.
*   The first character is `H`; its ASCII code is `72`.
*   The _current value_ increases to `72`.
*   The _current value_ is multiplied by `17` to become `1224`.
*   The _current value_ becomes `_200_` (the remainder of `1224` divided by `256`).
*   The next character is `A`; its ASCII code is `65`.
*   The _current value_ increases to `265`.
*   The _current value_ is multiplied by `17` to become `4505`.
*   The _current value_ becomes `_153_` (the remainder of `4505` divided by `256`).
*   The next character is `S`; its ASCII code is `83`.
*   The _current value_ increases to `236`.
*   The _current value_ is multiplied by `17` to become `4012`.
*   The _current value_ becomes `_172_` (the remainder of `4012` divided by `256`).
*   The next character is `H`; its ASCII code is `72`.
*   The _current value_ increases to `244`.
*   The _current value_ is multiplied by `17` to become `4148`.
*   The _current value_ becomes `_52_` (the remainder of `4148` divided by `256`).

So, the result of running the HASH algorithm on the string `HASH` is `_52_`.

The _initialization sequence_ (your puzzle input) is a comma-separated list of steps to start the Lava Production Facility. _Ignore newline characters_ when parsing the initialization sequence. To verify that your HASH algorithm is working, the book offers the sum of the result of running the HASH algorithm on each step in the initialization sequence.

For example:

    rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7

This initialization sequence specifies 11 individual steps; the result of running the HASH algorithm on each of the steps is as follows:

*   `rn=1` becomes `_30_`.
*   `cm-` becomes `_253_`.
*   `qp=3` becomes `_97_`.
*   `cm=2` becomes `_47_`.
*   `qp-` becomes `_14_`.
*   `pc=4` becomes `_180_`.
*   `ot=9` becomes `_9_`.
*   `ab=5` becomes `_197_`.
*   `pc-` becomes `_48_`.
*   `pc=6` becomes `_214_`.
*   `ot=7` becomes `_231_`.

In this example, the sum of these results is `_1320_`. Unfortunately, the reindeer has stolen the page containing the expected verification number and is currently running around the facility with it excitedly.

Run the HASH algorithm on each step in the initialization sequence. _What is the sum of the results?_ (The initialization sequence is one long line; be careful when copy-pasting it.)

To begin, [get your puzzle input](15/input).

Answer:  

You can also \[Shareon [Twitter](https://twitter.com/intent/tweet?text=%22Lens+Library%22+%2D+Day+15+%2D+Advent+of+Code+2023&url=https%3A%2F%2Fadventofcode%2Ecom%2F2023%2Fday%2F15&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)\] this puzzle.

(function(i,s,o,g,r,a,m){i\['GoogleAnalyticsObject'\]=r;i\[r\]=i\[r\]||function(){ (i\[r\].q=i\[r\].q||\[\]).push(arguments)},i\[r\].l=1\*new Date();a=s.createElement(o), m=s.getElementsByTagName(o)\[0\];a.async=1;a.src=g;m.parentNode.insertBefore(a,m) })(window,document,'script','//www.google-analytics.com/analytics.js','ga'); ga('create', 'UA-69522494-1', 'auto'); ga('set', 'anonymizeIp', true); ga('send', 'pageview');