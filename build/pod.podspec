Pod::Spec.new do |spec|
  spec.name         = 'Gchain'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/chainstone/go-chainstone'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Chainstoneeum Client'
  spec.source       = { :git => 'https://github.com/chainstone/go-chainstone.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gchain.framework'

	spec.prepare_command = <<-CMD
    curl https://gchainstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gchain.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
