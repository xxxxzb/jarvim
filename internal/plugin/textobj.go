// Package plugin provides ...
package plugin

const DeinSandWich = `
[[plugins]]
repo = 'machakann/vim-sandwich'
on_map = { vonx = '<Plug>(operator-sandwich-' }
hook_add = '''
    let g:sandwich_no_default_key_mappings = 1
    let g:operator_sandwich_no_default_key_mappings = 1
    let g:textobj_sandwich_no_default_key_mappings = 1
'''
`

const DeinTextObj = `
[[plugins]]
repo = 'kana/vim-operator-user'

[[plugins]]
repo = 'kana/vim-operator-replace'
on_map = { vnx = '<Plug>' }

[[plugins]]
repo = 'kana/vim-textobj-user'

[[plugins]]
repo = 'terryma/vim-expand-region'
on_map = { x = '<Plug>' }

[[plugins]]
repo = 'AndrewRadev/splitjoin.vim'
on_map = { n = '<Plug>Splitjoin' }

[[plugins]]
repo = 'AndrewRadev/dsf.vim'
on_map = { n = '<Plug>Dsf' }
hook_add = '''
	let g:dsf_no_mappings = 1
'''

[[plugins]]
repo = 'kana/vim-niceblock'
on_map = { x = '<Plug>' }
hook_add = '''
	let g:niceblock_no_default_key_mappings = 0
'''
[[plugins]]
repo = 'osyo-manga/vim-textobj-multiblock'
on_map = { ox = '<Plug>' }
hook_add = '''
	let g:textobj_multiblock_no_default_key_mappings = 1
'''

[[plugins]]
repo = 'kana/vim-textobj-function'
on_map = { ox = '<Plug>' }
hook_add = '''
	let g:textobj_function_no_default_key_mappings = 1
'''
`
