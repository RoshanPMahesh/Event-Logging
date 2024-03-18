import pandas as pand
import matplotlib.pyplot as matplot

dataframes = pand.read_csv("delay.csv", header=None) # theres no header in the csv

minimum = []
maximum = []
median_num = []
ninety_perc = []

for i in range(2, 103):
    # row = dataframes[(dataframes[0] >= i) & (dataframes[0] <= i + 1)] # selects the row based on the column -> represented as an array with whatevers at the row
    row = dataframes[(dataframes[0] == i)]
    if len(row) == 0:
        minimum.insert(len(minimum), 0)
        maximum.insert(len(maximum), 0)
        median_num.insert(len(median_num), 0)
        ninety_perc.insert(len(ninety_perc), 0)
    else:
        minimum.insert(len(minimum), row[1].min())
        maximum.insert(len(maximum), row[1].max())
        median_num.insert(len(median_num), row[1].median())
        ninety_perc.insert(len(ninety_perc), row[1].quantile(0.9))

matplot.plot(range(0, 101), minimum, label="Minimum Delay")
matplot.plot(range(0, 101), maximum, label="Maximum Delay")
matplot.plot(range(0, 101), median_num, label="Median Delay")
matplot.plot(range(0, 101), ninety_perc, label="90th Percentile Delay")

matplot.title("Delay over Time")
matplot.ylabel("Logging Delay in Seconds")
matplot.xlabel("Time in Seconds")
matplot.legend(loc=1)
# matplot.ylim(-0.001, 0.007) # to zoom into plots with spikes in delay (8 nodes)

matplot.show()