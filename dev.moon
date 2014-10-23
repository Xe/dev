#!/usr/bin/env moon

-- Copyright (C) 2014 Sam Dodrill <xena@yolo-swag.com> All rights reserved.
--
-- This software is provided 'as-is', without any express or implied
-- warranty. In no event will the authors be held liable for any damages
-- arising from the use of this software.
--
-- Permission is granted to anyone to use this software for any purpose,
-- including commercial applications, and to alter it and redistribute it
-- freely, subject to the following restrictions:
--
-- 1. The origin of this software must not be misrepresented; you must not
--    claim that you wrote the original software. If you use this software
--    in a product, an acknowledgment in the product documentation would be
--    appreciated but is not required.
--
-- 2. Altered source versions must be plainly marked as such, and must not be
--    misrepresented as being the original software.
--
-- 3. This notice may not be removed or altered from any source
--    distribution.

export yaml = require "yaml"

--- yaml2Table loads a yaml file as text and returns it as a file.
--  Returns a table representing the yaml document and an error string or nil.
yaml2Table = (fname) -> --> table, error
  fin = io.open fname, "r"
  if fin == nil
    return {}, "File read (#{fname}) failed."

  data = fin\read "*all"
  ret = yaml.load data

  fin\close!

  return ret, nil

--- doCommand returns the output and return status of a command.
doCommand = (command) -> --> string, number
  n = os.tmpname!
  code = os.execute command .. " > " .. n
  lines = {}

  for line in io.lines n
    table.insert lines, line

  os.remove n

  lines, code

--- pairsByKeys is an iterator for a table aplhabetically
--  adapted from here: http://www.lua.org/pil/19.3.html
pairsByKeys = (t) -> --> function
  a = {}
  for n in pairs t
    table.insert a, n

  table.sort a

  i = 0
  iter = ->
    i += 1
    if a[i] == nil
      return nil
    else
      return a[i], t[a[i]]
  iter

export commands = {
  up: {"Brings up a development container", ->
    dcommand = "docker run -idt --name #{data.projname}-dev --hostname dev:#{data.projname} "
    path = "/home/#{data.user}/dev/"
    localdir = os.getenv "PWD"

    -- Is this go?
    if data.golang
      path = "/home/#{data.user}/go/src/#{repopath}"

    -- Add source code directory to mount
    dcommand ..= "-v #{localdir}:#{path} "

    -- Add ssh keys if needed
    if data.ssh
      dcommand ..= "-v /home/#{os.getenv "USER"}/.ssh:/home/#{data.user}/.ssh "

    image = data.base
    if data.base == nil and data.overlay ~= nil
      image = "dev-#{data.projname}"

    -- Append image name
    dcommand ..= image

    print "Starting up container for #{data.projname}"

    -- Start up the docker container
    lines, status = doCommand dcommand
    if status ~= 0
      print "Launch failed. Does this container already exist?"
      for _,line in pairs lines
        print "docker: " .. line
      os.exit status

    ctid = lines[#lines]

    -- report to user
    print "#{data.projname}-dev (#{ctid\sub(1,6)}) running!"
    print "To use this container please attach to it with:"
    print "    $ docker attach #{data.projname}-dev"
  }

  down: { "Destroys a development container", ->
    lines, status = doCommand "docker rm -f #{data.projname}-dev"
    if status ~= 0
      os.exit status

    print "Container destroyed."
  }

  establish: { "Create a Docker image from the manifest", ->
    if data.overlay == nil
      print "Error: docker container overlay not made. Exiting."
      os.exit 1

    dir = os.tmpname!
    os.remove dir

    out, err = doCommand "mkdir #{dir}"
    if err ~= 0
      print "Could not make temporary directory #{dir}. Dying."
      os.exit err

    fout = io.open "#{dir}/Dockerfile", "w"
    fout\write data.overlay
    fout\close!

    proc = io.popen "docker build -t dev-#{data.projname} #{dir}"
    for line in proc\lines!
      print line

    proc\close!

    _, err = doCommand "rm -rf #{dir}"
    if err ~= 0
      os.exit err

    print "Docker image dev-#{data.projname} created."
  }
}

if #arg == 0
  print "dev version 0.2\n"

  print "Usage: dev [command] <manifest>\n"

  print "  if manifest is undefined the default value"
  print "  .dev.yaml will be used.\n"

  print "Available commands:"
  for name,cmd in pairsByKeys commands
    print "%12s   %s"\format(name, cmd[1])

  os.exit 1

command = arg[1]\lower!

if commands[command] ~= nil
  export fname = arg[2]
  if fname == nil
    fname = ".dev.yaml"

  export data, err = yaml2Table fname
  if err ~= nil
    print "Cannot open file #{fname}. Please make sure this file exists or is readable."
    os.exit 1

  commands[command][2]!
