#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/udp.h>
#include <bpf/bpf_helpers.h>

int five(int a,int b) {
	return 5+a + b;
}


void main() {
	int a = 10;
int b = five(a,a);;
int c = a + b;
}

