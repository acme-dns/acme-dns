package nameserver

import (
	"testing"

	"github.com/miekg/dns"
)

func TestNameserver_isOwnChallenge(t *testing.T) {
	type fields struct {
		OwnDomain string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "is own challenge",
			fields: fields{
				OwnDomain: "some-domain.test.",
			},
			args: args{
				name: "_acme-challenge.some-domain.test.",
			},
			want: true,
		},
		{
			name: "challenge but not for us",
			fields: fields{
				OwnDomain: "some-domain.test.",
			},
			args: args{
				name: "_acme-challenge.some-other-domain.test.",
			},
			want: false,
		},
		{
			name: "not a challenge",
			fields: fields{
				OwnDomain: "domain.test.",
			},
			args: args{
				name: "domain.test.",
			},
			want: false,
		},
		{
			name: "other request challenge",
			fields: fields{
				OwnDomain: "domain.test.",
			},
			args: args{
				name: "my-domain.test.",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Nameserver{
				OwnDomain: tt.fields.OwnDomain,
			}
			if got := n.isOwnChallenge(tt.args.name); got != tt.want {
				t.Errorf("isOwnChallenge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameserver_isAuthoritative(t *testing.T) {
	type fields struct {
		OwnDomain string
		Domains   map[string]Records
	}
	type args struct {
		q dns.Question
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "is authoritative own domain",
			fields: fields{
				OwnDomain: "auth.domain.",
			},
			args: args{
				q: dns.Question{Name: "auth.domain."},
			},
			want: true,
		},
		{
			name: "is authoritative other domain",
			fields: fields{
				OwnDomain: "auth.domain.",
				Domains: map[string]Records{
					"other-domain.test.": {Records: nil},
				},
			},
			args: args{
				q: dns.Question{Name: "other-domain.test."},
			},
			want: true,
		},
		{
			name: "is authoritative sub domain",
			fields: fields{
				OwnDomain: "auth.domain.",
				Domains: map[string]Records{
					"other-domain.test.": {Records: nil},
				},
			},
			args: args{
				q: dns.Question{Name: "sub.auth.domain."},
			},
			want: true,
		},
		{
			name: "is not authoritative own",
			fields: fields{
				OwnDomain: "auth.domain.",
				Domains: map[string]Records{
					"other-domain.test.": {Records: nil},
				},
			},
			args: args{
				q: dns.Question{Name: "special-auth.domain."},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Nameserver{
				OwnDomain: tt.fields.OwnDomain,
				Domains:   tt.fields.Domains,
			}
			if got := n.isAuthoritative(tt.args.q.Name); got != tt.want {
				t.Errorf("isAuthoritative() = %v, want %v", got, tt.want)
			}
		})
	}
}
