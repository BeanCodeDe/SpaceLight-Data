import csv
import sys

class Room():
    def __init__(self, x,y):
       self.pos = [[x,y]]

print("Please input ship name:")

shipName = input()


with open('ship_type/'+shipName+'/room.csv', newline='') as csvfile:
    spamreader = csv.reader(csvfile, delimiter=',', quotechar=' ')
    elementList = []
    for row in spamreader:
        elementRow = []
        for element in row:
            length = len(element)
            if element != '0' and length != 4:
                print("Element has not length of 4: "+ element)
            elementRow.append(element)
        elementList.append(elementRow)

print("parsed csv")

rooms = []
y = 0
for row in elementList:
    x = 0
    for element in row:
        if (element[0]=='D' or element[0]=='W') and (element[3]=='D' or element[3]=='W'):
            room = Room(x,y)
            rooms.append(room)
        elif element != '0':
            for room in rooms:
                toAddPos = []
                for pos in room.pos:
                    if element[0]== '-' and pos[0]==x and pos[1]==y-1:
                        toAddPos.append([x,y])
                        break
                    elif element[2]== '-' and pos[0]==x-1 and  pos[1]==y:
                        toAddPos.append([x,y])
                        break
                room.pos += toAddPos
        x += 1
    y +=1

print("rooms created")

roomString = "["
for room in rooms:
    minX = 999
    minY = 999
    maxX = -1
    maxY = -1
    for pos in room.pos:
        if minX > pos[0]:
            minX = pos[0]
        if maxX < pos[0]:
            maxX = pos[0]
        if minY > pos[1]:
            minY = pos[1]
        if maxY < pos[1]:
            maxY = pos[1]

    roomString += "{"
    roomString += "\"PosX\":"+str(minX)+","
    roomString += "\"PosY\":"+str(minY)+","
    roomString += "\"SizeX\":"+str(maxX-minX)+","
    roomString += "\"SizeY\":"+str(maxY-minY)
    roomString += "},"

roomString += "]"
print(roomString)
