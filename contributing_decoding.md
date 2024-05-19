# Contributing Vehicle Decoding

Please have vehicle decoding contributions be in the following format:

Below is an example with a single entry, for multiple signals just include more entries.
```json
[
  {
    "name": "vehicleSpeed",
    "comment": "front left wheel speed, kmh",
    "header": "7df",
    "mode": "22",
    "pid": "1135",
    "formula": "python:bytes_to_int(messages[0].data[-2:]) * 0.05625",
    "interval_seconds": 10,
    "protocol": "6",
    "can_flow_control_clear": true,
    "can_flow_control_id_pair": "713,77D",
    "can_flow_control_filter": "607,FFF",
    "can_extended_address": "07"
  }
]
```

### Explanation

If you're using an autopi to decode, with eg. the `obd.query` [command](https://docs.autopi.io/core/commands/core-commands-obd/#obdquery), 
you will notice we're just parameterizing the requests into a json format.

- name: can be anything, as long as it makes sense. We'll convert it to the right thing in our [model-garage COVESA spec](https://github.com/DIMO-Network/model-garage)
- comment: any useful info, the units eg. kmh, degrees Celcius
- header,mode,pid: bytes in HEX format (not decimal format)
- can_flow: all optional, just depending if the vehicle needs them to be decoded. 
- interval_seconds: optional, we'll most likely determine how often this signal should be queried
- formula: how to decode the incoming bytes. 
  - Prepend with `python:` if it is just using the built-in autopi formula.
  - Prepend with `dbc:` (prefered, but may not be possible for all scenarios) if you can use the standard dbc formula to decode the bytes. [tooling here](https://www.csselectronics.com/pages/dbc-editor-can-bus-database)
example: `24|16@1+ (0.125,0) [0|8031.875] "rpm"`
  - Prepend with `custom:` if this needs special logic or just easier to explain in words, eg. "may return 1 or 2 frames, single byte after the PID convert to decimal". We'll figure out code to process.