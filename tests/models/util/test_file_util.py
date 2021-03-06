import os
from unittest import TestCase

from ezored.models.util.file_util import FileUtil
from testfixtures import tempdir


class TestFileUtil(TestCase):
    @tempdir()
    def test_create_dir(self, d):
        os.chdir(d.path)

        dir_name_1 = 'new-dir-1'
        dir_name_2 = 'new-dir-2'

        FileUtil.create_dir(dir_name_1)
        FileUtil.create_dir(os.path.join(dir_name_1, dir_name_2))

        self.assertTrue(os.path.isdir(dir_name_1))
        self.assertTrue(os.path.isdir(os.path.join(dir_name_1, dir_name_2)))

    @tempdir()
    def test_remove_dir(self, d):
        os.chdir(d.path)

        dir_name_1 = 'new-dir-1'
        dir_name_2 = 'new-dir-2'

        FileUtil.create_dir(dir_name_1)
        FileUtil.create_dir(os.path.join(dir_name_1, dir_name_2))
        FileUtil.remove_dir(dir_name_1)

        self.assertFalse(os.path.isdir(dir_name_1))
        self.assertFalse(os.path.isdir(os.path.join(dir_name_1, dir_name_2)))

    @tempdir()
    def test_remove_file(self, d):
        os.chdir(d.path)

        filename = 'new-file.txt'

        FileUtil.write_to_file('.', filename, 'content test')

        self.assertTrue(os.path.isfile(filename))

        FileUtil.remove_file(filename)

        self.assertFalse(os.path.isfile(filename))

    @tempdir()
    def test_write_to_file(self, d):
        os.chdir(d.path)

        filename = 'new-file.txt'

        FileUtil.write_to_file('.', filename, 'content test')

        self.assertTrue(os.path.isfile(filename))
        self.assertEqual(os.path.getsize(filename), 12)
