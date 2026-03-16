# ezbpf

## Overview
The ezbpf project is an advanced implementation of the eBPF (Extended Berkeley Packet Filter) functionality, providing developers with a powerful toolset for network diagnostics, security and various performance analysis applications.

## Language Features
- **Flexibility**: Allows developers to write code that can be executed safely in the kernel, enabling on-the-fly modifications to kernel behavior without the need for full reboots.
- **Performance**: Offers high performance by running bytecode in a restricted environment with a focus on speed and efficiency.
- **Safety**: Ensures safety by performing extensive checks on the code written, preventing misuse or crashes in the kernel space.

## Installation
To install ezbpf, you can clone the repository and build the project:
```bash
git clone https://github.com/toastsandwich/ezbpf.git
cd ezbpf
make
```

## Example Code
Below is an example of using ezbpf to monitor network traffic:

```c
#include <ezbpf.h>

int filter_func(struct __sk_buff *skb) {
    // Access packet data and filter accordingly
    if (skb->protocol == htons(ETH_P_IP)) {
        return EZBPF_PASS; // Allow IP packets
    }
    return EZBPF_DROP; // Drop all other packets
}

int main() {
    // Setup BPF program and attach to relevant hook
    ezbpf_attach(filter_func);
    return 0;
}
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for more information on how to get started.