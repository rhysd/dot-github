guard :shell do
  watch /\.go$/ do |m|
    puts "\033[93m#{Time.now}: #{File.basename m[0]}\033[0m"
    system "go build"
  end
end
