package id

import (
	"sync"
	"testing"
)

func Test_uuidUsecase_GenerateID(t *testing.T) {
	type fields struct {
		id int64
	}
	tests := []struct {
		name        string
		fields      fields
		numSubTasks int
		want        string
	}{
		{
			name:        "TestConcurrent",
			fields:      fields{},
			numSubTasks: 10000,
			want:        "10001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &uuidUsecase{
				id: tt.fields.id,
			}
			var wg sync.WaitGroup
			wg.Add(tt.numSubTasks)
			for i := 0; i < tt.numSubTasks; i++ {
				go func() {
					defer wg.Done()
					u.GenerateID()
				}()
			}
			wg.Wait()
			if got := u.GenerateID(); got != tt.want {
				t.Errorf("uuidUsecase.GenerateID() = %v, want %v", got, tt.want)
			}
		})
	}
}
