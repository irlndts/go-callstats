# go-callstats
Test application for callstats.io
Problem description bellow.

The algorithm is to use two heaps to keep lower and higher halfs of metrics separately. The median is the lowest value of the higher heap or (higher value of lower + the lowest value of the higher heap)/2
And also there is a queue to keep data in slider range.

Test 1 is implemented via TestGetMedian in median_test.go

### Complexity
#### Time complexity:
At worst, there are three heap insertions and two heap deletions from the top for min-max heaps + deletion after slider is full. After Each of these takes about O(log(n)) time. Finding the mean takes constant O(1) time since the tops of heaps are directly accessible.
We also need to insert and delete from queue - it's O(2*log(n)).
So the result is o(8*log(n) + O(1)) ≈ `O(log(n))`

#### Space
`O(2 * size of slider)` - we keep data in queue and in heaps separately, and we don't need to keep more than slider requeres. It's linear.

### What to improve?
- more tests! Cover heaps etc.
- validate input parameters
- try to use benchamars and go tools like `go tool pprof` to check problems with memory and cpu.

=======================

# Original Problem
Conference calls are happening over the Internet and participants are reporting metrics from the endpoints to callstats.io in real-time. The conference call duration can be anything from a few minutes long (industry average for call centers) to a few hours (typical for team meetings) or even longer. In some cases the conference calls last for days (between sites for example). Network engineers rely on particular metrics and look for specific statistical trends to identify calls with bad quality and furthermore, to make decisions. For example, a network engineer might want to identify a bad call based on common statistical measurements: average, max, min, median, standard deviation, skew, etc.

For this problem set, let's assume that we get just one metric, the network delay between participants, reported at regular intervals, typically once every 100ms. At the end of a 5 minute call there would be: 3000 measurements, in 30 minutes: 18000, and in 6 hours: 216000 measurements.

We need to calculate the median over a set of measurements for providing reliable information about the network delay. The measurements are stored in a sliding window, which limits the number of items and the intervals between the first and last element in it. For an instance, a sliding window can contains maximum 5 received items and the maximum interval is 25s. 

Your task is to implement a sliding window contains limited amount of items and provides addDelay and getMedian interfaces. The former interface adds a delay value to-, the latter interface returns the median of the delays calculated over the items from the sliding window. 
You must limit your sliding window regarding to the number of items it can contain.


## Test 1:
An example is given below, using a sliding window with length of 3.

The delay measurement arrive one-by-one (iteration) in the following order:
100, 102, 101, 110, 120, 115,

The sliding window should look like this at each iteration:
100
100, 102,
100, 102, 101,
102, 101,  110,
101, 110,  120,
110,  120,  115,

Output: after each iteration (use \r\n delimiter)
-1
101
101
102
110
115

Help:
If only one element available in the sliding window the answer is -1.
If n is odd then Median (M) = value of ((n + 1)/2)th item from a sorted array of length n.
If n is even then Median (M) = value of [((n)/2)th item term + ((n)/2 + 1)th item term ] /2
