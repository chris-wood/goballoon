from matplotlib import cm
from mpl_toolkits.mplot3d import Axes3D
import matplotlib.pyplot as plt
import numpy as np
import sys

x = []
y = []
time_z = []
mem_z = []

with open(sys.argv[1], 'r') as fh:
    for line in fh:
        line = line.split(",")
        xval = int(line[0])
        yval = int(line[1])
        time = float(line[2])
        mem = int(line[4])

        x.append(xval)
        y.append(yval)
        time_z.append(time)
        mem_z.append(mem)

#x = [1000,1000,1000,1000,1000,5000,5000,5000,5000,5000,10000,10000,10000,10000,10000]
#y = [13,21,29,37,45,13,21,29,37,45,13,21,29,37,45]
#z = [75.2,79.21,80.02,81.2,81.62,84.79,87.38,87.9,88.54,88.56,88.34,89.66,90.11,90.79,90.87]

fig = plt.figure()
ax = fig.gca(projection='3d')
ax.plot_trisurf(x, y, time_z, cmap=cm.jet, linewidth=0.2)
ax.set_xlabel('Space [blocks]')
ax.set_ylabel('Iterations')
#ax.set_zlabel('Running Time [ms]')
ax.set_title("Running Time [ms]")
#plt.show()
plt.savefig("running_time.png")


fig = plt.figure()
ax = fig.gca(projection='3d')
ax.plot_trisurf(x, y, mem_z, cmap=cm.jet, linewidth=0.2)
ax.set_xlabel('Space [blocks]')
ax.set_ylabel('Iterations')
#ax.set_zlabel('Allocated Space [B]')
ax.set_title("Allocated Space [B]")
#plt.show()
plt.savefig("allocated_space.png")
