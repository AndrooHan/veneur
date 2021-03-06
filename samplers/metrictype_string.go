// Code generated by "stringer -type MetricType ./samplers"; DO NOT EDIT.

package samplers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CounterMetric-0]
	_ = x[GaugeMetric-1]
	_ = x[StatusMetric-2]
}

const _MetricType_name = "CounterMetricGaugeMetricStatusMetric"

var _MetricType_index = [...]uint8{0, 13, 24, 36}

func (i MetricType) String() string {
	if i < 0 || i >= MetricType(len(_MetricType_index)-1) {
		return "MetricType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MetricType_name[_MetricType_index[i]:_MetricType_index[i+1]]
}
