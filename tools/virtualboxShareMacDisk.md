ref:
https://blog.csdn.net/a962804835/article/details/72820355


firstly, create a shared file for virtualbox and macbook file.

```shell
cd /Applications/VirtualBox.app/Contents/Resources/VirtualBoxVM.app/Contents/MacOS

cp VBoxGuestAdditions.iso /Users/zidif/Downloads/

# using virtualbox cdrom to insert VBoxGuestAdditions.iso
mount /dev/cdrom /cdrom

cd /cdrom

sh ./VBoxLinuxAdditions.run
# install any necessary softwares...

mount -t vboxsf gopath /root/mac-gopath/

```
