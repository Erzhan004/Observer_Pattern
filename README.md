# Observer_Pattern

Teacher Interface
The Teacher interface defines the methods a teacher should have:
AddGroup(group Group): Adds a group to the teacher's list.
NotifyAll(message string): Notifies all groups with a message.

Group Interface
The Group interface defines the methods a group should have:

ReactToMessage(message string): Reacts to a message received by the group.
Teacher Struct
The teacher struct represents a teacher with a list of groups.

Group Struct
The group struct represents a group with a unique ID.

Functions
GetNewTeacher() teacher: Creates and returns a new teacher instance.
GetNewGroup(Id string) group: Creates and returns a new group instance with the given ID.

Main Function
In the main function, we create a teacher (teach1) and three groups (group1, group2, group3). The teacher adds these groups to its list and then notifies all groups with a message.
