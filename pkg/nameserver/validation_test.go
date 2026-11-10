package nameserver

import (
	"reflect"
	"testing"

	"github.com/miekg/dns"
)

func TestNameserver_answerOwnChallenge(t *testing.T) {
	type fields struct {
		personalAuthKey string
	}
	type args struct {
		q dns.Question
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dns.RR
		wantErr bool
	}{
		{
			name: "answer own challenge",
			fields: fields{
				personalAuthKey: "some key text",
			},
			args: args{
				q: dns.Question{
					Name:   "something",
					Qtype:  0,
					Qclass: 0,
				},
			},
			want: []dns.RR{
				&dns.TXT{
					Hdr: dns.RR_Header{Name: "something", Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 1},
					Txt: []string{"some key text"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Nameserver{}

			n.SetOwnAuthKey(tt.fields.personalAuthKey)
			if n.personalAuthKey != tt.fields.personalAuthKey {
				t.Errorf("failed to set personal auth key: got = %s, want %s", n.personalAuthKey, tt.fields.personalAuthKey)
				return
			}

			got, err := n.answerOwnChallenge(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("answerOwnChallenge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerOwnChallenge() got = %v, want %v", got, tt.want)
			}
		})
	}
}
