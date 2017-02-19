use strict;
use warnings;

# ls dist | xargs -I{} github-release upload --user t-mrt --repo gocha --tag v0.0.1 --name {} --file "dist/{}"

my $DIST_DIR = "dist";
my $BIN_NAME = "gocha";

open my $fh, '<', './cmd/gocha/version.go'
    or die "failed to open: $!";
my $v = 'v';
while (my $line = <$fh>) {
    if ($line =~ /version.*"(.*)"/i) {
        $v .= $1;
    }
}

my $os_archs = [
    {os => "darwin",  arch => "amd64"},
    {os => "darwin",  arch => "386"  },
    {os => "freebsd", arch => "amd64"},
    {os => "freebsd", arch => "386"  },
    {os => "freebsd", arch => "arm"  },
    {os => "linux",   arch => "amd64"},
    {os => "linux",   arch => "386"  },
    {os => "linux",   arch => "arm"  },
    {os => "netbsd",  arch => "amd64"},
    {os => "netbsd",  arch => "386"  },
    {os => "netbsd",  arch => "arm"  },
    {os => "openbsd", arch => "amd64"},
    {os => "openbsd", arch => "386"  },
    {os => "plan9",   arch => "amd64"},
    {os => "plan9",   arch => "386"  },
    {os => "windows", arch => "amd64"},
    {os => "windows", arch => "386"  },
    {os => "nacl",    arch => "386"  },
];

for (@$os_archs) {
    my $goos = $_->{os};
    my $goarch = $_->{arch};
    my $zip_name = $BIN_NAME . "_" . $v . "_" . $goos . "_" . $goarch . ".zip";

    system "GOOS=$goos GOARCH=$goarch go build -o $DIST_DIR/$BIN_NAME ./cmd/gocha";
    system "rm -f $DIST_DIR/$zip_name";
    system "zip -m -q $DIST_DIR/$zip_name $DIST_DIR/$BIN_NAME";

    print "Write: ${DIST_DIR}/${zip_name}\n";
}
