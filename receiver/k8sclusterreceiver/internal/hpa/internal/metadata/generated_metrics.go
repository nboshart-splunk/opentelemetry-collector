// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

type metricK8sHpaCurrentReplicas struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.hpa.current_replicas metric with initial data.
func (m *metricK8sHpaCurrentReplicas) init() {
	m.data.SetName("k8s.hpa.current_replicas")
	m.data.SetDescription("Current number of pod replicas managed by this autoscaler.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sHpaCurrentReplicas) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sHpaCurrentReplicas) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sHpaCurrentReplicas) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sHpaCurrentReplicas(cfg MetricConfig) metricK8sHpaCurrentReplicas {
	m := metricK8sHpaCurrentReplicas{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sHpaDesiredReplicas struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.hpa.desired_replicas metric with initial data.
func (m *metricK8sHpaDesiredReplicas) init() {
	m.data.SetName("k8s.hpa.desired_replicas")
	m.data.SetDescription("Desired number of pod replicas managed by this autoscaler.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sHpaDesiredReplicas) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sHpaDesiredReplicas) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sHpaDesiredReplicas) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sHpaDesiredReplicas(cfg MetricConfig) metricK8sHpaDesiredReplicas {
	m := metricK8sHpaDesiredReplicas{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sHpaMaxReplicas struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.hpa.max_replicas metric with initial data.
func (m *metricK8sHpaMaxReplicas) init() {
	m.data.SetName("k8s.hpa.max_replicas")
	m.data.SetDescription("Maximum number of replicas to which the autoscaler can scale up.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sHpaMaxReplicas) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sHpaMaxReplicas) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sHpaMaxReplicas) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sHpaMaxReplicas(cfg MetricConfig) metricK8sHpaMaxReplicas {
	m := metricK8sHpaMaxReplicas{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sHpaMinReplicas struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.hpa.min_replicas metric with initial data.
func (m *metricK8sHpaMinReplicas) init() {
	m.data.SetName("k8s.hpa.min_replicas")
	m.data.SetDescription("Minimum number of replicas to which the autoscaler can scale up.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sHpaMinReplicas) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sHpaMinReplicas) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sHpaMinReplicas) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sHpaMinReplicas(cfg MetricConfig) metricK8sHpaMinReplicas {
	m := metricK8sHpaMinReplicas{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	startTime                   pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity             int                 // maximum observed number of metrics per resource.
	metricsBuffer               pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                   component.BuildInfo // contains version information
	metricK8sHpaCurrentReplicas metricK8sHpaCurrentReplicas
	metricK8sHpaDesiredReplicas metricK8sHpaDesiredReplicas
	metricK8sHpaMaxReplicas     metricK8sHpaMaxReplicas
	metricK8sHpaMinReplicas     metricK8sHpaMinReplicas
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                   pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:               pmetric.NewMetrics(),
		buildInfo:                   settings.BuildInfo,
		metricK8sHpaCurrentReplicas: newMetricK8sHpaCurrentReplicas(mbc.Metrics.K8sHpaCurrentReplicas),
		metricK8sHpaDesiredReplicas: newMetricK8sHpaDesiredReplicas(mbc.Metrics.K8sHpaDesiredReplicas),
		metricK8sHpaMaxReplicas:     newMetricK8sHpaMaxReplicas(mbc.Metrics.K8sHpaMaxReplicas),
		metricK8sHpaMinReplicas:     newMetricK8sHpaMinReplicas(mbc.Metrics.K8sHpaMinReplicas),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/k8sclusterreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricK8sHpaCurrentReplicas.emit(ils.Metrics())
	mb.metricK8sHpaDesiredReplicas.emit(ils.Metrics())
	mb.metricK8sHpaMaxReplicas.emit(ils.Metrics())
	mb.metricK8sHpaMinReplicas.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordK8sHpaCurrentReplicasDataPoint adds a data point to k8s.hpa.current_replicas metric.
func (mb *MetricsBuilder) RecordK8sHpaCurrentReplicasDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sHpaCurrentReplicas.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sHpaDesiredReplicasDataPoint adds a data point to k8s.hpa.desired_replicas metric.
func (mb *MetricsBuilder) RecordK8sHpaDesiredReplicasDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sHpaDesiredReplicas.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sHpaMaxReplicasDataPoint adds a data point to k8s.hpa.max_replicas metric.
func (mb *MetricsBuilder) RecordK8sHpaMaxReplicasDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sHpaMaxReplicas.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sHpaMinReplicasDataPoint adds a data point to k8s.hpa.min_replicas metric.
func (mb *MetricsBuilder) RecordK8sHpaMinReplicasDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sHpaMinReplicas.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
