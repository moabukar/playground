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

- Let's say you have LAN 1 and LAN 2 which are isolated networks. Devices on each local network can communicate with each other but not outside of that layer 2 network
  - Ethernet is a layer 2 protocol. Generally used for local networks. Long distance point to point connections are not possible with ethernet and will use more suitable protocols like PPP (Point to Point Protocol), MPLS (Multi Protocol Label Switching), Frame Relay, ATM (Asynchronous Transfer Mode)
  - Layer 2 is the layer of the OSI stack which moves frames. Moving frames from a local source to a local destination.
  - So to move data between different local networks, which is known as inter-networking, this is where the internet comes from, we need layer 3
- Layer 3 adds the internet protcol or IP. You get IP addresses which across networking addresses which you can assign to devices and these can be used to communicate between different local networks using routing.
- IP packets are moved source to destination across the internet through many intermediate networks. Devices called routers which are layer 3 devices are used to move IP packets across different networks.
- They encapsulate IP packets into layer 2 ethernet frames. So they take the IP packet and put it into a layer 2 frame. 
- Encapsulaton here means that an IP packet is put inside an ethernet frame for that part of the journey.

### IP & Packets

- Packets in many ways are similar to frames. They contain data to be moved and they have a source and dest address. With frames, both the source and dest address are moved across a LAN. With IP packets, the source and dest address are moved across the internet and could be on opposite sides of the planet.

### Packet structure of IPv4 vs IPv6

#### IPv4

- Every packet has a source and destination IP address
- Protocol field (which is layer 4) like ICMP, TCP, UDP
  - If you're storing TCP data inside a packet, this value will be 6
  - If you're storing UDP data inside a packet, this value will be 17
  - If you're storing ICMP data inside a packet, this value will be 1
- Bulk of the field within a packet is taken up by the data. 
- A field called TTL
- And other stuff too..

#### IPv6

- Source & IP (Bigger & therefore, larger addresses)
- Data
- Hop limit (like TTL)
- And other stuff...
  
### IP addressing (v4)

- Example: 133.33.3.7
  - 133.33 is the 'network' part
  - 3.7 is the 'host' part which represents hosts on that network
  - In this case, 3.7 is laptop on the network 133.33
- Note: If the network part of the IP address match between 2 different IP addresses, they are on the same IP network. If they don't match, they are on different IP networks.
- Each part is 8 bits. So 4 parts = 32 bits. This is why IPv4 is 32 bits. Each 8 bit part is called an octet.
- So for example: 133.33.3.7 and 133.33.33.37
  - They are on the same network because the first 2 octets match
  - They are different hosts because the last 2 octets are different
- Now IP addresses are either statically assigned by huamns which is known as static IP addressing or they are assigned dynamically/automatically by a service called DHCP (Dynamic Host Configuration Protocol). So the servers on your network running DHCP server software will automatically assign IP addresses to devices on your network. This is known as dynamic IP addressing.

### Subnet Masks

- A subnet mask is configured on a host device in addition to an IP address e.g. 255.255.0.0 & this is the same as a /16 prefix.
- Subnet masks are used to identify the network & host part of an IP address.
- 


## Route Tables & Routes 

- Example of packet moving between routers
  - Home > ISP >  AWS/Upstream ISP/Netflix (Remote networks)
  - Create a packet on our local device which has our IP and the source IP address
  - Default route 0.0.0.0/0 sends all packets to ISP
  - Your packet that was generated is now in the router that has multiple network interface cards connecting to all of the remote networks (AWS/Netflix/Upstream ISP)
  - The ISP uses route tables to forward packets/data to remote networks. Every router will have at least 1 route table
  - A route table is a collection of routes: Destination to Next hop/target
  - Packets are routed, hop by hop across the internet. From source to destination. Router compares packet destination IP and route table for matching destinations. The more specific prefixes are preferred (0 lowest, 32 higher). Packet is forwarded to next hop/target. This process is repeated until the packet reaches the destination.
  - Routers can be statically populate or there are protocols such as BGP (border gateway protocol) which allow routers to communicate with each other to exchange which networks they know about and this is how the core of the internet functions
  - Important note: when our ISP router is forwarding the packet through to the AWS router, it's forwarding at layer 2
  - It wraps the packet in a frame. The packet doesn't change. The frame though, it has the AWS router's MAC address as its destination. But how do we determine the MAC address of the AWS router here?
  - We use something called the ARP (Address Resolution Protocol) Continued below!

## ARP (Address Resolution Protocol)

- The ARP is used when you have a layer 3 packet and you want to encapsulate it inside a frame and send that frame to a MAC address. We don't initially know the MAC address and we need a protocol which can find the MAC address for a given IP address
- For example if you communicate with AWS, AWS will be the destination of the IP packets but we will forward via our home router which is the default gateway and so we will need the MAC address of that default gateway to send the frame to containing the packet. This is where ARP comes in
- ARP will give you the MAC address for a given IP address


Example of ARP:

- 2 laptops: Laptop A (133.33.3.7) wants to send data to Laptop B (133.33.3.10). THey are both on the same network since the IPs have same subnet masks
- Laptop A takes the data and passes it to layer 3 which creates packet. The packet has its IP address as the source and laptop B as the destination IP
- Now we need a way of being able to create a frame to put that packet in for a transmission. We need the MAC address of laptop B
- This is what ARP does for us. It’s a process which runs between layer 2 and layer 3. Because both IPs have same subnet mask and it knows they are on the same local network, this is a direct connection. Routers aren’t required here. We don’t use routers for this type of comms
- ARP broadcasts on layer 2. It sends an ARP frame to all Fs as a MAC address and it asks who has the IP address 133.33.3.10. Laptop B is also running ARP. The ARP software sees this broadcast and it responds by saying im that IP address and here is my MAC address
- Now Laptop A has the destination MAC address of laptop B, it can use it to build a frame, encapsulate the packet in this frame and once frame is ready, it can be given to layer 1 and sent across the physical network. Layer 1 receives the physical raw bit stream and hands it off to layer 2 of the laptop. Layer 2 software of the laptop B reviews the dest MAC address and knows it meant for it.
- So it strips off the frame and sends the packet to its layer 3 software. Layer 3 sees the packet and sees it is the intended destination and it de-encapsulates the data and hands it to the game.

## Layer 3 - IP routing

- Example of a packet moving from multiple routers
- ARP is used to obtain a MAC address of another target. If a packet is going from another network, it uses a router. If it’s leaving a router to another one. ARP is used to find the MAC address of the default gateway of the router. The router picks it up in a frame and strips it but realises it’s not meant for it. It encapsulates it ina frame and knows the next hop for another router. The next router picks it up and strips it but finds the destination on the same network. So it uses ARP to find the MAC address of the destination. The destination picks it up and takes the data.

## Layer 3 Summary

- IP addresses (IPv4/6) - cross network addressing
- ARP - find the MAC address for a certain IP
- Routes: where to forward packets
- Route tables: containing multiple routes
- Router: moves packets from source to destination (encapsulating in L2 frames on the way)
- This allows for device to device comms over the internet
- IP doesn’t provide method for channels of comms (SRC IP to DST IP only)
- Layer 3 provides packets and packets only have source and dest IP. So if we have 2 devices, you can only have one stream of comms. So you can’t have different apps on the devices communicating at the same time. And this limitation is solved by layer 4.
- In theory, packets can be delivered out of order. No guarantee it will take the same route


# Layer 4 - Transport

- 
