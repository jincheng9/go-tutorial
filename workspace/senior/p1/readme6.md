# Mac系统查看Go开发相关的系统设置

## 查看CPU信息

```bash
MacBook-Air:go-tutorial $ sysctl -a machdep.cpu
machdep.cpu.max_basic: 20
machdep.cpu.max_ext: 2147483656
machdep.cpu.vendor: GenuineIntel
machdep.cpu.brand_string: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
machdep.cpu.family: 6
machdep.cpu.model: 61
machdep.cpu.extmodel: 3
machdep.cpu.extfamily: 0
machdep.cpu.stepping: 4
machdep.cpu.feature_bits: 9221959987971750911
machdep.cpu.leaf7_feature_bits: 35399595 0
machdep.cpu.leaf7_feature_bits_edx: 2617247232
machdep.cpu.extfeature_bits: 1241984796928
machdep.cpu.signature: 198356
machdep.cpu.brand: 0
machdep.cpu.features: FPU VME DE PSE TSC MSR PAE MCE CX8 APIC SEP MTRR PGE MCA CMOV PAT PSE36 CLFSH DS ACPI MMX FXSR SSE SSE2 SS HTT TM PBE SSE3 PCLMULQDQ DTES64 MON DSCPL VMX EST TM2 SSSE3 FMA CX16 TPR PDCM SSE4.1 SSE4.2 x2APIC MOVBE POPCNT AES PCID XSAVE OSXSAVE SEGLIM64 TSCTMR AVX1.0 RDRAND F16C
machdep.cpu.leaf7_features: RDWRFSGS TSC_THREAD_OFFSET BMI1 AVX2 SMEP BMI2 ERMS INVPCID FPU_CSDS RDSEED ADX SMAP IPT MDCLEAR IBRS STIBP L1DF SSBD
machdep.cpu.extfeatures: SYSCALL XD 1GBPAGE EM64T LAHF LZCNT PREFETCHW RDTSCP TSCI
machdep.cpu.logical_per_package: 16
machdep.cpu.cores_per_package: 8
machdep.cpu.microcode_version: 47
machdep.cpu.processor_flag: 6
machdep.cpu.mwait.linesize_min: 64
machdep.cpu.mwait.linesize_max: 64
machdep.cpu.mwait.extensions: 3
machdep.cpu.mwait.sub_Cstates: 286531872
machdep.cpu.thermal.sensor: 1
machdep.cpu.thermal.dynamic_acceleration: 1
machdep.cpu.thermal.invariant_APIC_timer: 1
machdep.cpu.thermal.thresholds: 2
machdep.cpu.thermal.ACNT_MCNT: 1
machdep.cpu.thermal.core_power_limits: 1
machdep.cpu.thermal.fine_grain_clock_mod: 1
machdep.cpu.thermal.package_thermal_intr: 1
machdep.cpu.thermal.hardware_feedback: 0
machdep.cpu.thermal.energy_policy: 1
machdep.cpu.xsave.extended_state: 7 832 832 0
machdep.cpu.xsave.extended_state1: 1 0 0 0
machdep.cpu.arch_perf.version: 3
machdep.cpu.arch_perf.number: 4
machdep.cpu.arch_perf.width: 48
machdep.cpu.arch_perf.events_number: 7
machdep.cpu.arch_perf.events: 0
machdep.cpu.arch_perf.fixed_number: 3
machdep.cpu.arch_perf.fixed_width: 48
machdep.cpu.cache.linesize: 64
machdep.cpu.cache.L2_associativity: 8
machdep.cpu.cache.size: 256
machdep.cpu.tlb.inst.large: 8
machdep.cpu.tlb.data.small: 64
machdep.cpu.tlb.data.small_level1: 64
machdep.cpu.address_bits.physical: 39
machdep.cpu.address_bits.virtual: 48
machdep.cpu.core_count: 2
machdep.cpu.thread_count: 4
machdep.cpu.tsc_ccc.numerator: 0
machdep.cpu.tsc_ccc.denominator: 0
```

倒数第4行的`machdep.cpu.core_count: 2` 表示有2个核心

倒数第3行的`machdep.cpu.thread_count: 4`表示有4个线程，也就是4个逻辑CPU

这台机器是1个物理CPU(physical cpu)，每个物理CPU上有2个核心(2 cores per physical cpu)，每个核心是超线程，总共4个逻辑CPU(logical cpu)

也可以点击电脑左上角`苹果图标->关于本机->系统报告`，查看具体信息。

```markdown
硬件概览：

  型号名称：	MacBook Air
  型号标识符：	MacBookAir7,2
  处理器名称：	Dual-Core Intel Core i5
  处理器速度：	1.6 GHz
  处理器数目：	1
  核总数：	2
  L2缓存（每个核）：	256 KB
  L3缓存：	3 MB
  超线程技术：	已启用
  内存：	8 GB
  Boot ROM版本：	427.0.0.0.0
  SMC版本（系统）：	2.27f2
  序列号（系统）：	FVFSC135H3QF
  硬件UUID：	3242BDAE-C5CD-5162-AEC3-7001C7C451C5
```



## Go语言里获取逻辑CPU个数以及GOMAXPROCS

```go
// cpu.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// nCPU is the number of logical cpu
	nCPU := runtime.NumCPU()
	// num is the number of currrent GOMAXPROCS
	// default is the value of runtime.NumCPU()
	num := runtime.GOMAXPROCS(0)
	// 4 4
	fmt.Println(num, nCPU)
}
```

