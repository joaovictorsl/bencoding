This is a simple module for encoding/decoding to/from **Bencode**.

**Bencode** is the encoding used by the peer-to-peer file sharing system **BitTorrent** for storing and transmitting loosely structured data.
It is most commonly used in torrent files, which are simply bencoded dictionaries.

Supported types:

- byte strings
- integers
- lists
- dictionaries
