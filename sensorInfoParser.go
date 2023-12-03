package main

import (
	"regexp"
	"strings"
)

func parseTemperatureData(input string) map[string]interface{} {
	temperatureData := make(map[string]interface{})

	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`([a-zA-Z_]+):\s+([-+]?\d*\.\d+|\d+)\s+(C)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 4 {
			entryName := match[1]
			value := match[2]
			unit := match[3]

			entry := map[string]interface{}{
				"value": value,
				"unit":  unit,
			}

			temperatureData[entryName] = entry
		}
	}

	return temperatureData
}

func get_temperatures() map[string]interface{} {
	input := `
	iio_hwmon_ams_ctrl-isa-0000
	Adapter: ISA adapter
	vcc_pspll:       +1.19 V  
	vcc_psbatt:      +0.00 V  
	vccint:          +0.84 V  
	vccbram:         +0.84 V  
	vccaux:          +1.79 V  
	vcc_psddr_pll:   +1.81 V  
	vccpsintfp_ddr:  +0.84 V  
	vccint:          +0.84 V  
	vccaux:          +1.78 V  
	vccvrefp:        +0.42 V  
	vccvrefn:        +0.00 V  
	vccbram:         +0.84 V  
	vccplintlp:      +0.84 V  
	vccplintfp:      +0.84 V  
	vccplaux:        +1.80 V  
	vccams:          +0.60 V  
	vccpsintlp:      +0.84 V  
	vccpsintfp:      +0.84 V  
	vccpsaux:        +1.80 V  
	vccpsddr:        +1.10 V  
	vccpsio3:        +1.80 V  
	vccpsio0:        +1.80 V  
	vccpsio1:        +1.80 V  
	vccpsio2:        +1.80 V  
	psmgtravcc:      +0.84 V  
	psmgtravtt:      +1.79 V  
	vccams:          +0.60 V  
	ps_temp:         +48.6 C  
	remote_temp:     +51.6 C  
	pl_temp:         +51.1 C  
	
	 
	
	iio_hwmon_adc1-isa-0000
	Adapter: ISA adapter
	emmc0_icc:    +0.00 V  
	emmc0_iccq:   +0.00 V  
	emmc1_icc:    +0.00 V  
	emmc1_iccq:   +0.00 V  
	
	 
	
	iio_hwmon_pim_adc-isa-0000
	Adapter: ISA adapter
	ADC Channel 0:  +0.01 V  
	ADC Channel 1:  +0.00 V  
	ADC Channel 2:  +0.00 V  
	ADC Channel 3:  +0.00 V  
	ADC Channel 4:  +0.00 V  
	ADC Channel 5:  +0.00 V  
	ADC Channel 6:  +0.00 V  
	ADC Channel 7:  +0.00 V  
	
	 
	
	iio_hwmon_adc2-isa-0000
	Adapter: ISA adapter
	cur_vcc_1v8_aux:  +0.21 V  
	vcc_0v85:         +0.84 V  
	vcc_1v1:          +1.10 V  
	vcc_1v8:          +1.81 V  
	
	 
	
	iio_hwmon_adc0-isa-0000
	Adapter: ISA adapter
	cur_vcc_in:   +0.39 V  
	vcc_in:      +11.46 V  
	vcc_3v3:      +3.33 V  
	vcc_5v0:      +4.97 V 
`
	// replace the hard coded input with the output of the command
	parsedData := parseTemperatureData(input)
	// fmt.Printf("%#v\n", parsedData)
	return parsedData

}
