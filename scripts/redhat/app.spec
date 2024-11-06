Name: codegen
Version: %VERSION%
Release: 0
License: MIT
Requires: openssl
Prefix: /usr
Summary: Code generator - cli tool
Group: Applications/System
BuildArch: x86_64

%description
Code generator - cli tool

%pre

%prep

%build

%install
mkdir -p %{buildroot}/usr/bin
cp -rfa %{_app_dir}/* %{buildroot}/usr/bin

%post -p /bin/bash
chown -R root:root /usr/bin/codegen
chmod -R 777 /usr/bin/codegen

if [ -f "/usr/lib/systemd/system/codegen.service" ]; then
    rm -rf /usr/lib/systemd/system/codegen.service
    systemctl disable codegen.service
    systemctl stop codegen.service
    systemctl daemon-reload
fi

cat <<EOF > /etc/systemd/system/codegen.service
[Unit]
Description=Code generator - cli tool

[Service]
Type=simple
WorkingDirectory=/usr/bin
ExecStart=/usr/bin/codegen
Restart=always
RestartSec=10
KillSignal=SIGINT
SyslogIdentifier=codegen
User=root
Group=root

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable codegen.service
systemctl restart codegen.service

%clean

%files
%defattr(0770, root, root)
/usr/bin/codegen

%define _unpackaged_files_terminate_build 0
