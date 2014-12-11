#! -*- coding: utf-8 -*-

from distutils.core import setup
from distutils.extension import Extension
from Cython.Distutils import build_ext

ext_modules = [
    Extension("main", ["main.py"]),
]

setup(
    name = "bench",
    cmdclass = {"build_ext" : build_ext},
    ext_modules = ext_modules,
)
