# Importing the modules
import shutil


def backup():
    # src_dir = os.getcwd()  # get pwd
    src_dir1 = 'E:/World of Warcraft/_retail_/Interface'
    src_dir2 = 'E:/World of Warcraft/_retail_/WTF'

    dest_dir1 = "D:/Backups/Wowui backup/Interface"
    dest_dir2 = "D:/Backups/Wowui backup/WTF"

    try:
        shutil.rmtree(dest_dir1)
        shutil.rmtree(dest_dir2)
        shutil.copytree(src_dir1, dest_dir1, symlinks=False, ignore=None, copy_function=shutil.copy2, ignore_dangling_symlinks=False)
        shutil.copytree(src_dir2, dest_dir2, symlinks=False, ignore=None, copy_function=shutil.copy2, ignore_dangling_symlinks=False)
    except:
        print('An error has occurred')
    else:
        print('Complete without errors')


def restore():
    # src_dir = os.getcwd()  # get pwd
    dest_dir1 = 'E:/World of Warcraft/_retail_/Interface'
    dest_dir2 = 'E:/World of Warcraft/_retail_/WTF'

    src_dir1 = "D:/Backups/Wowui backup/Interface"
    src_dir2 = "D:/Backups/Wowui backup/WTF"

    try:
        shutil.rmtree(dest_dir1)
        shutil.rmtree(dest_dir2)
        shutil.copytree(src_dir1, dest_dir1, symlinks=False, ignore=None, copy_function=shutil.copy2,
                        ignore_dangling_symlinks=False)
        shutil.copytree(src_dir2, dest_dir2, symlinks=False, ignore=None, copy_function=shutil.copy2,
                        ignore_dangling_symlinks=False)
    except:
        print('An error has occurred')
    else:
        print('Complete without errors')


def main():
    print('backup or restore?')
    inp = input('>')

    if inp.lower() == 'backup':
        backup()
    elif inp.lower() == 'restore':
        restore()
    else:
        main()


main()
