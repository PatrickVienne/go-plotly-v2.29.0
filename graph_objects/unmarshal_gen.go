package grob

import (
	"encoding/json"
	"errors"
)

type unmarshalType struct {
	Type TraceType `json:"type,omitempty"`
}

// UnmarshalTrace decodes an array of bytes into a Trace interface.
func UnmarshalTrace(data []byte) (Trace,error) {
	traceType := unmarshalType{}
	err := json.Unmarshal(data, &traceType)
	if err != nil {
		return nil, err
	}
	switch traceType.Type {
    case TraceTypeArea:
        trace := &Area{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeBar:
        trace := &Bar{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeBarpolar:
        trace := &Barpolar{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeBox:
        trace := &Box{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeCandlestick:
        trace := &Candlestick{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeCarpet:
        trace := &Carpet{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeChoropleth:
        trace := &Choropleth{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeChoroplethmapbox:
        trace := &Choroplethmapbox{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeCone:
        trace := &Cone{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeContour:
        trace := &Contour{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeContourcarpet:
        trace := &Contourcarpet{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeDensitymapbox:
        trace := &Densitymapbox{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeFunnel:
        trace := &Funnel{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeFunnelarea:
        trace := &Funnelarea{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeHeatmap:
        trace := &Heatmap{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeHeatmapgl:
        trace := &Heatmapgl{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeHistogram:
        trace := &Histogram{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeHistogram2d:
        trace := &Histogram2d{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeHistogram2dcontour:
        trace := &Histogram2dcontour{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeImage:
        trace := &Image{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeIndicator:
        trace := &Indicator{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeIsosurface:
        trace := &Isosurface{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeMesh3d:
        trace := &Mesh3d{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeOhlc:
        trace := &Ohlc{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeParcats:
        trace := &Parcats{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeParcoords:
        trace := &Parcoords{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypePie:
        trace := &Pie{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypePointcloud:
        trace := &Pointcloud{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeSankey:
        trace := &Sankey{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScatter:
        trace := &Scatter{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScatter3d:
        trace := &Scatter3d{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScattercarpet:
        trace := &Scattercarpet{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScattergeo:
        trace := &Scattergeo{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScattergl:
        trace := &Scattergl{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScattermapbox:
        trace := &Scattermapbox{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScatterpolar:
        trace := &Scatterpolar{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScatterpolargl:
        trace := &Scatterpolargl{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeScatterternary:
        trace := &Scatterternary{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeSplom:
        trace := &Splom{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeStreamtube:
        trace := &Streamtube{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeSunburst:
        trace := &Sunburst{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeSurface:
        trace := &Surface{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeTable:
        trace := &Table{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeTreemap:
        trace := &Treemap{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeViolin:
        trace := &Violin{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeVolume:
        trace := &Volume{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    case TraceTypeWaterfall:
        trace := &Waterfall{}
        err = json.Unmarshal(data,trace)
        if err != nil {
            return nil, err
        }
        return trace, nil
    default:
        return nil, errors.New("Trace Type is not registered")
	}
}