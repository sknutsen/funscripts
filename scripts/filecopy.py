# Importing the modules
import shutil

# src_dir = os.getcwd()  # get the current working dir
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
