package builder

import (
	"reflect"
	"testing"
)

func TestResourcePoolConfigBuilder_Build(t *testing.T) {
	tests := []struct {
		name    string
		fields  *ResourcePoolConfigBuilder
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			fields: &ResourcePoolConfigBuilder{
				name:     "",
				maxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			fields: &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 0,
				maxIdle:  10,
				minIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: &ResourcePoolConfigBuilder{
				name: "test",
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: 0,
				maxIdle:  0,
				minIdle:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ResourcePoolConfigBuilder{
				name:     tt.fields.name,
				maxTotal: tt.fields.maxTotal,
				maxIdle:  tt.fields.maxIdle,
				minIdle:  tt.fields.minIdle,
			}
			got, err := b.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourcePoolConfigBuilder.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourcePoolConfigBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
