Name: codegen
Version: %VERSION%
Release: 0
License: MIT
Requires: openssl
Prefix: /opt
Summary: Code generator - cli tool
Group: Applications/System
BuildArch: x86_64

%description
Code generator - cli tool

%pre

%prep

%build

%install
mkdir -p %{buildroot}/opt/codegen
mkdir -p %{buildroot}/usr/bin
cp -rfa %{_app_dir}/* %{buildroot}/opt/codegen
ln -s /opt/codegen/codegen %{buildroot}/usr/bin/codegen

%post -p /bin/bash
chown -R root:root /opt/codegen
chmod -R 770 /opt/codegen

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
WorkingDirectory=/opt/codegen
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
/opt/codegen/*
/usr/bin/codegen

%define _unpackaged_files_terminate_build 0
