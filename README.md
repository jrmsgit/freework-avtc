Tasks
=====

1. Take this line of code as an example:

{"level":"debug","ts":1578025817,"conn_id":5406983,"state":"closed","Tx":404,"Rx":879}

2. Write logs generator that will generate logs based on the given format that will contain 24 hours data from Feb 20th, 2020. The expected outcome is the log file generated as fast as possible and containing 1 record per each second, conn_id by 1 should be incremented for each rеcord, with random values for TX and RX from 0 - 10000.

3. Write simple prometheus exporter that will parse log file and export into file two metrics:

3.1 histogram of TX values with buckets:
    < 500
    < 1000
    < 5000
    < 9000
    < inf

3.1 histogram of RX values with buckets:
    < 500
    < 1000
    < 5000
    < 9000
    < inf

ts - unix time
Please also measure the execution time for generation and parsing (separately).

Implementation
==============

2. run ./gen.sh script to generate ./out/log.txt file with random data.

3. run ./export.sh to populate prometheus histograms.

Run ./bench.sh for benchmarks on generating the log's data and parsing it.
