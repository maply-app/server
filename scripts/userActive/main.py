# import requests
# import time
#
#
# class Config:
#     serverAddress = "http://127.0.0.1"
#     requestsTimeout = 3
#     userToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY4OTU5MjUsImlh\
#     dCI6MTY3NDMwMzkyNSwic3ViIjoiYTZkNmUyNzktNjA1ZS00YjU1LWE4ZjEtNzQxNjRjNTcwMTY5In0\
#     .P8rNp7xrD-igTO3nak4J2nPZte7kbX86a32sn7TWhdk"
#     startCoords = [55.701879, 37.930528]
#     # 55.701879, 37.930528
#     # 55.701848, 37.930580
#     updateRange = []
#
#
# class Test:
#     def __init__(self, conf: Config):
#         self.serverAddress = conf.serverAddress
#         self.requestsTimeout = conf.requestsTimeout
#         self.userToken = conf.userToken
#
#     def updateCoords(self):
#         ...
#
#     def sendRequest(self):
#         ...
#
#     def start(self):
#         while True:
#             self.sendRequest()
#             time.sleep(self.requestsTimeout)
#

import math

c1 = [55.701472, 37.938312]
c2  = [55.702325, 37.938859]

# https://www.movable-type.co.uk/scripts/latlong.html

# >> Distance –> meters
R = 6367444.6571225
f1 = c1[0] * math.pi / 180
f2 = c2[0] * math.pi / 180
deltaF = (c2[0] - c1[0]) * math.pi / 180
deltaY = (c2[1] - c1[1]) * math.pi / 180
a = math.sin(deltaF / 2) * math.sin(deltaF / 2) \
    + math.cos(f1) * math.cos(f2) \
    * math.sin(deltaY / 2) * math.sin(deltaY/2)
c = 2 * math.atan2(math.sqrt(a), math.sqrt(1 - a))
d = R * c

print(d) # 100.79... meters

# >> Speed –> km/h
d = 100.79556014059384
t = 35 # time sec
s = d / t / 1000 * 3600
print(s) # 10.36... km/h

#

