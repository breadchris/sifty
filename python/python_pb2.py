# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: python.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cpython.proto\x12\x06python\"#\n\x10SummarizeRequest\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\"$\n\x11SummarizeResponse\x12\x0f\n\x07summary\x18\x01 \x01(\t\"!\n\x11TranscribeRequest\x12\x0c\n\x04\x66ile\x18\x01 \x01(\t\"+\n\x12TranscribeResponse\x12\x15\n\rtranscription\x18\x01 \x01(\t2\x8f\x01\n\x06Python\x12\x43\n\nTranscribe\x12\x19.python.TranscribeRequest\x1a\x1a.python.TranscribeResponse\x12@\n\tSummarize\x12\x18.python.SummarizeRequest\x1a\x19.python.SummarizeResponseB\x07Z\x05./genb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'python_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\005./gen'
  _SUMMARIZEREQUEST._serialized_start=24
  _SUMMARIZEREQUEST._serialized_end=59
  _SUMMARIZERESPONSE._serialized_start=61
  _SUMMARIZERESPONSE._serialized_end=97
  _TRANSCRIBEREQUEST._serialized_start=99
  _TRANSCRIBEREQUEST._serialized_end=132
  _TRANSCRIBERESPONSE._serialized_start=134
  _TRANSCRIBERESPONSE._serialized_end=177
  _PYTHON._serialized_start=180
  _PYTHON._serialized_end=323
# @@protoc_insertion_point(module_scope)