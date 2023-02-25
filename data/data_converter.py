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
    errorFile = False
    with open(shipName+'/ship.csv', newline='') as csvfile:
        spamreader = csv.reader(csvfile, delimiter=',', quotechar=' ')
        elementList = []
        for row in spamreader:
            elementRow = []
            for element in row:
                length = len(element)
                if element != '0' and length != 4:
                    print("Element has not length of 4: "+ element)
                    errorFile = True
                elementRow.append(element)
            elementList.append(elementRow)
    if errorFile:
        sys.exit()
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

def parseRoomListToRoomJson(roomPlaceList: List[RoomPlace]) -> str:
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

def searchRoom(oomPlaceList: List[RoomPlace], x: int, y: int) -> RoomBlock:
    for roomPlace in roomPlaceList:
        for roomBlock in roomPlace.roomBlockList:
            if roomBlock.x == x and roomBlock.y == y:
                return roomBlock
    return None

def parseRoomPlaceListToDoors(roomPlaceList: List[RoomPlace])-> List[Door]:
    print("parseRoomPlaceListToDoors")
    doorList = []
    for roomPlace in roomPlaceList:
        for roomBlock in roomPlace.roomBlockList:
            if roomBlock.roomDesc[1] == 'D':
                foundRoomBlock = searchRoom(roomPlaceList, roomBlock.x+1, roomBlock.y)
                if foundRoomBlock == None:
                    print("No matching door found from room ("+str(roomBlock.x)+" | "+str(roomBlock.y)+") with description "+ roomBlock.roomDesc+" to room ("+str(roomBlock.x+1)+" | "+str(roomBlock.y)+")")
                    sys.exit()
                door = Door(roomBlock.id, foundRoomBlock.id)
                doorList.append(door)
            if roomBlock.roomDesc[2] == 'D':
                foundRoomBlock = searchRoom(roomPlaceList, roomBlock.x, roomBlock.y+1)
                if foundRoomBlock == None:
                    print("No matching door found from room ("+str(roomBlock.x)+" | "+str(roomBlock.y)+") with description "+ roomBlock.roomDesc+" to room ("+str(roomBlock.x)+" | "+str(roomBlock.y+1)+")")
                    sys.exit()
                door = Door(roomBlock.id, foundRoomBlock.id)
                doorList.append(door)
    return doorList


def parseRoomListToDoorJson(doorList: List[Door]) -> str:
    print("parseRoomListToDoorJson")
    doorJson = "["
    for door in doorList:
        doorJson += "{"
        doorJson += "\"RoomBlockOneId\":"+str(door.roomBlockOneId)+","
        doorJson += "\"RoomBlockTwoId\":"+str(door.roomBlockTwoId)
        doorJson += "},"
    doorJson = doorJson[:-1]
    doorJson += "]"
    return doorJson
    
print("Please input ship name:")
#shipName = input()
shipName = "spacelight"

elementList = parseFileToElementList(shipName)
roomPlaceList = parseElementsListToRooms(elementList)
roomJson = parseRoomListToRoomJson(roomPlaceList)

print("ROOM STRING:")
print(roomJson)

doorList = parseRoomPlaceListToDoors(roomPlaceList)
doorJson = parseRoomListToDoorJson(doorList)

print("DOOR STRING:")
print(doorJson)