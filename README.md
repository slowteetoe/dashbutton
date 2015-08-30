Dash
===
Inspired by <a href="https://medium.com/@edwardbenson/how-i-hacked-amazon-s-5-wifi-button-to-track-baby-data-794214b0bdd8">this article</a> but since I have an Intel Edison lying around, and I can compile a binary for the Edison... I'm doing it in Go.

Eventually (or maybe not that long from now) this will allow me to trigger events in the real world, basically anything that I can hook up to an Edison (which means all Arduino/Seed/etc boards).

First Time Setup
---
Follow the instructions that came with your Dash button, but in the final step, do NOT pick the actual product you want to replenish.  Just close the shopping app instead.
Run the util/identify.go script - you'll most likely (definitely?) need to run it as root: sudo GOPATH=/your/go/path go run identify.go -inf=eth0 (you can use the inf flag to specify which interface you want to listen on)
Press the Dash button - you should see something like:
SourceMAC[de:ad:be:ef:01:02]
Save this MAC address, it's how your Dash button will be identified.

Usage
---
(TBD) 
