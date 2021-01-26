make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

autorun_version=`./bin/autorun -v`
echo "build version: $autorun_version"

# cross_compiles
make -f ./Makefile.cross-compiles

rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        autorun_dir_name="autorun_${autorun_version}_${os}_${arch}"
        autorun_path="./packages/autorun_${autorun_version}_${os}_${arch}"

        echo $autorun_dir_name

        if [ "x${os}" = x"windows" ]; then
            if [ ! -f "./autorun_${os}_${arch}.exe" ]; then
                continue
            fi
            if [ ! -f "./autorun_${os}_${arch}.exe" ]; then
                continue
            fi
            mkdir ${autorun_path}
            mv ./autorun_${os}_${arch}.exe ${autorun_path}/autorun.exe
        else
            if [ ! -f "./autorun_${os}_${arch}" ]; then
                continue
            fi
            if [ ! -f "./autorun_${os}_${arch}" ]; then
                continue
            fi
            mkdir ${autorun_path}
            mv ./autorun_${os}_${arch} ${autorun_path}/autorun
        fi  

        cp ../LICENSE ${autorun_path}

        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${autorun_dir_name}.zip ${autorun_dir_name}
        else
            tar -zcf ${autorun_dir_name}.tar.gz ${autorun_dir_name}
        fi  
        cd ..
        rm -rf ${autorun_path}
    done
done

cd -