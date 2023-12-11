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
- 
