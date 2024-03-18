import pandas as pand
import matplotlib.pyplot as matplot

dataframes = pand.read_csv("bandwidth.csv", header=None) # theres no header in the csv

average = []
rate = 1000

for i in range(101):
    row = dataframes[(dataframes[0] == i)]
    if len(row) == 0:
        average.append(0)
    else:
        average.append(row[1].mean()/rate)

matplot.plot(range(0, 101), average, label="Average")

matplot.title("Bandwidth over Time")
matplot.ylabel("Bandwidth in Kbps") # change based on rate
matplot.xlabel("Time in Seconds")
matplot.legend()

matplot.show()