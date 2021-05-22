package promo_user_has_voucher

import (
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func Test_repoImpl_FindById(t *testing.T) {

	type args struct {
		user_id string
		id      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "",
			args:    args{
				user_id: "",
				id:      "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repoImpl{
				collection: tt.fields.collection,
			}
			got, err := r.FindById(tt.args.user_id, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
