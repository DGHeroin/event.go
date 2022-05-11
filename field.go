package event
type (
    F map[string]interface{}
)
func (f F) Int(key string) int {
    if p, ok := f[key]; ok {
        if val, ok := p.(int); ok {
            return val
        }
    }
    return 0
}
func (f F) Int64(key string) int64 {
    if p, ok := f[key]; ok {
        if val, ok := p.(int64); ok {
            return val
        }
    }
    return 0
}
func (f F) Float32(key string) float32 {
    if p, ok := f[key]; ok {
        if val, ok := p.(float32); ok {
            return val
        }
    }
    return 0
}
func (f F) Float64(key string) float64 {
    if p, ok := f[key]; ok {
        if val, ok := p.(float64); ok {
            return val
        }
    }
    return 0
}
func (f F) String(key string) string {
    if p, ok := f[key]; ok {
        if val, ok := p.(string); ok {
            return val
        }
    }
    return ""
}
func (f F) Bool(key string) bool {
    if p, ok := f[key]; ok {
        if val, ok := p.(bool); ok {
            return val
        }
    }
    return false
}
