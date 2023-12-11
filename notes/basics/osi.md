# OSI

## Layer 1 - Physical

- Layer 1 specifications define the transmission and reception of raw BIT STREAMS between a device and a shared physical medium.
- Physical medium can be copper (electrical signals), fibre (light), wifi (radio frequencies or waves)
- Like a physical hub used to connect multiple devices together
- No device addressing, all data is processed by all devices. It's liek shouting in a room without saying any names and everyone hears it. This is a limitation and it's solved by layer 2.
- No media access control

## Layer 2 - Data Link

- Runs over layer 1 and it requires a functional layer 1 to operate
- Higher layers build on lower layers
- Frames are a format used in layer 2 to send information over a layer 2 network
- Layer 2 also introduces a new unique hardware address aka a MAC address. This address is uniquely assigned to a specific hardware.
  - MAC address is formed of 2 parts:
    - OUI (Organizationally Unique Identifier) - first 3 bytes of the MAC address (assigned to companies who manufacture network devices)
    - NIC (Network Interface Controller) - last 3 bytes of the MAC address
- Layer 2 uses layer 1: This means that a layer 2 or ethernet frame can be transmitted onto the shared physical medium by layer 1. These are converted to voltages, RF or light
- Layer 2 provides frames & Layer 1 handles the physical transmission and reception onto and from the physical medium
- So when layer 1 is transmitting a frame onto the physical medium, layer 1 doesn't understand the frame. Layer 1 simply transmits raw data onto the physical medium. 

- Layer 2 has different parts:
  - Preamble bits & start frame delimiter (the function of this is to allow devices to know that it's the start of a frame)
  - Next is the destination and source MAC address. All devices on a layer 2 network have a unique MAC address
  - And a frame can be sent to a specific device by putting its MAC address destination. Or you can put all Fs if you want to send the frame to every device on the local network. This is called a broadcast.
  -  Next is Ethertype which is a layer 3 protocol. Layer 3 uses layer 2 frames for device to device communication on a local network. So when recieving a frame at the other side, you need to know which layer 3 protocol originally put data into that frame. A common example is IP or Internet Protocol. And this is what the ethertype field is for. It tells the receiving device which layer 3 protocol put data into the frame.
  -  These 3 fields (Dest MAC address, Source MAC address, Ethertype) are called the MAC header of the frame.
  -  After the header, it's the payload. It contains the data that the frame is sending. The data is generally provided by the layer 3 protocol. This process is called encapsulation. You have something which layer 3 generates, often this is an IP packet and it's put inside an ethernet frame. It's encapsulated in that frame. The frame delivers that data to a different layer 2 destination. 
  -  At the end of the frame is the frame check sequence which is used to identify any errors in the frame. 
- Using a HUB (let's say we have 4 devices connected to a HUB) - a hub is a layer 1 device. The data can have collisions. What you need is a switch. 
  - A switch is a layer 2 device. Works the same way physically as a HUB but it understands layer 2.
  - It maintains a MAC address table. Switches over tie learn what's connected to each port. When a switch sees frames, it can interpret frames, it can intercept them and see the source and destination MAC addresses. So over time, with this network, the MAC address table will be populated with each of the devices. So the switch will store the MAC addresses it sees on a port and the port itself. 
  - Switches are intelligent. They don't just repeat the physical level. They interpret the frames and they can make decisions based on the source and destination MAC address table.
  - So switches store and forward frames. It doesn't repeat like a dumb layer 1 device. It means it wont forward collisions. In fact, each port on the switch is a separate collision domain. So if there's a collision on one port, it won't affect the other ports. The switch will not forward that corrupted data to the other ports.
  - Layer 2 is the foundation for all networks which we use day to day. It's how our wired networks work. It's how our wifi networks work. It's how the internt works which is a huge collection of interconnected layer 2 networks.
  - The name itself stands for an inter-network of networks.

### Summary of when adding layer 2

- Identifiable devices using MAC addresses. Allows for device to device comms
- Media access control (sharing) - devices can share media in a nice way - avoiding collisions and cross talk
- Collision detection (when using switches)
- Unicase 1:1, Broadcast 1:All, Multicast 1:Many
- We have switches - basically like hubs but with super powers (layer 2) which are more intelligent and can make better decisions compared to layer 1 - ability to scale and avoid collisions

## Layer 3 - Network

- 
