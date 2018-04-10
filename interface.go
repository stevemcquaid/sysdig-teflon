type CustomMetricsProvider interface {
	// GetRootScopedMetricByName fetches a particular metric for a particular root-scoped object.
	GetRootScopedMetricByName(groupResource schema.GroupResource, name string, metricName string) (*custom_metrics.MetricValue, error)

	// GetRootScopedMetricByName fetches a particular metric for a set of root-scoped objects
	// matching the given label selector.
	GetRootScopedMetricBySelector(groupResource schema.GroupResource, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error)

	// GetNamespacedMetricByName fetches a particular metric for a particular namespaced object.
	GetNamespacedMetricByName(groupResource schema.GroupResource, namespace string, name string, metricName string) (*custom_metrics.MetricValue, error)

	// GetNamespacedMetricByName fetches a particular metric for a set of namespaced objects
	// matching the given label selector.
	GetNamespacedMetricBySelector(groupResource schema.GroupResource, namespace string, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error)

	// ListAllMetrics provides a list of all available metrics at
	// the current time.  Note that this is not allowed to return
	// an error, so it is reccomended that implementors cache and
	// periodically update this list, instead of querying every time.
	ListAllMetrics() []MetricInfo
}
