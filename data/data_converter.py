import csv
import sys
from typing import List

class RoomPlace():
    def __init__(self, id):
        self.id = id
        self.roomBlockList = []

    def addRoomBlock(self, roomBlock):
        self.roomBlockList.append(roomBlock)

    def getRoomId(self, x, y):
        for roomBlock in self.roomBlockList:
            if roomBlock.x == x and roomBlock.y == y:
                return roomBlock.id
        return -1

class RoomBlock():
    def __init__(self, id, x, y, roomDesc):
        self.id = id
        self.x = x
        self.y = y
        self.roomDesc = roomDesc

class Door():
    def __init__(self, roomBlockOneId, roomBlockTwoId):
        self.roomBlockOneId = roomBlockOneId
        self.roomBlockTwoId = roomBlockTwoId

def parseFileToElementList(shipName):
    print("parseFileToElementList")
    with open(shipName+'/ship.csv', newline='') as csvfile:
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
    return elementList

def parseElementsListToRooms(elementList)-> List[RoomPlace]:
    print("parseElementsListToRooms")
    roomPlaceList = []
    roomPlaceId = 0
    roomBlockId = 0
    y = 0
    for row in elementList:
        x = 0
        for element in row:
            if (element[0]=='D' or element[0]=='W') and (element[3]=='D' or element[3]=='W'):
                room = RoomPlace(roomPlaceId)
                roomPlaceId = 1 + roomPlaceId
                roomBlock = RoomBlock(roomBlockId,x,y,element)
                roomBlockId = 1 + roomBlockId
                room.addRoomBlock(roomBlock)
                roomPlaceList.append(room)
            elif element != '0':
                for room in roomPlaceList:
                    for roomBlock in room.roomBlockList:
                        if element[0]== '-' and roomBlock.x==x and roomBlock.y==y-1:
                            roomBlock = RoomBlock(roomBlockId,x,y,element)
                            roomBlockId = 1 + roomBlockId
                            room.addRoomBlock(roomBlock)
                            break
                        elif element[3]== '-' and roomBlock.x==x-1 and roomBlock.y==y:
                            roomBlock = RoomBlock(roomBlockId,x,y,element)
                            roomBlockId = 1 + roomBlockId
                            room.addRoomBlock(roomBlock)
                            break
            x += 1
        y +=1
    return roomPlaceList

def parseRoomListToRoomJson(roomPlaceList: RoomPlace) -> str:
    print("parseRoomListToRoomJson")
    roomJson = "["
    for roomPlace in roomPlaceList:
        roomJson += "{"
        roomJson += "\"Id\":"+str(roomPlace.id)+","
        roomJson += "\"RoomBlockList\": ["

        for roomBlock in roomPlace.roomBlockList:
            roomJson += "{"
            roomJson += "\"Id\":"+str(roomBlock.id)+","
            roomJson += "\"PosX\":"+str(roomBlock.x)+","
            roomJson += "\"PosY\":"+str(roomBlock.y)
            roomJson += "},"
    
        roomJson = roomJson[:-1]
        roomJson += "]"
        roomJson += "},"
    roomJson = roomJson[:-1]
    roomJson += "]"
    return roomJson
    
print("Please input ship name:")
#shipName = input()
shipName = "spacelight"

elementList = parseFileToElementList(shipName)
roomPlaceList = parseElementsListToRooms(elementList)
roomJson = parseRoomListToRoomJson(roomPlaceList)

print("ROOM STRING:")
print(roomJson)
