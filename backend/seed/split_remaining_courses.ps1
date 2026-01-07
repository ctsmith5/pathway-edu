$ErrorActionPreference = 'Stop'

$seedDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$seedPath = Join-Path $seedDir 'seed.go'

if (-not (Test-Path -LiteralPath $seedPath)) {
  throw "seed.go not found at: $seedPath"
}

$lines = Get-Content -LiteralPath $seedPath

function Find-LineIndex([string]$needle) {
  $m = ($lines | Select-String -SimpleMatch $needle | Select-Object -First 1)
  if (-not $m) { return $null }
  return ($m.LineNumber - 1)
}

function Extract-ModuleRangeById([string]$moduleId) {
  $idNeedle = 'ID:    "' + $moduleId + '"'
  $idIdx = Find-LineIndex $idNeedle
  if ($null -eq $idIdx) { throw "Could not find module ID: $moduleId" }

  $start = $null
  for ($i = $idIdx; $i -ge 0; $i--) {
    if ($lines[$i] -match '^\t\t\t\t\{\s*$') { $start = $i; break }
  }
  if ($null -eq $start) { throw "Could not find module start for: $moduleId" }

  $end = $null
  for ($i = $idIdx; $i -lt $lines.Count; $i++) {
    if ($lines[$i] -match '^\t\t\t\t\},\s*$') { $end = $i; break }
  }
  if ($null -eq $end) { throw "Could not find module end for: $moduleId" }

  return @{ start = $start; end = $end }
}

function Write-ModuleFile([string]$moduleId, [string]$funcName, [string]$fileName) {
  $range = Extract-ModuleRangeById $moduleId
  $block = @($lines[$range.start..$range.end])

  # Replace leading `{` and trailing `},` with a module return block
  $block[0] = "`treturn models.Module{"
  $block[$block.Count - 1] = "`t}"

  $out = @(
    'package seed',
    '',
    'import "github.com/pathway/backend/models"',
    '',
    ('func ' + $funcName + '() models.Module {')
  )
  $out += $block
  $out += @('}', '')

  Set-Content -LiteralPath (Join-Path $seedDir $fileName) -Value $out -Encoding utf8
}

function Get-CourseDescription([string]$courseTitle) {
  $titleNeedle = 'Title:       "' + $courseTitle + '"'
  $titleIdx = Find-LineIndex $titleNeedle
  if ($null -eq $titleIdx) { throw "Could not find course title: $courseTitle" }

  for ($i = $titleIdx + 1; $i -lt [Math]::Min($titleIdx + 10, $lines.Count); $i++) {
    if ($lines[$i] -match 'Description:\s+"([^"]+)"') {
      return $Matches[1]
    }
  }
  throw "Could not extract Description for course: $courseTitle"
}

function Write-CourseFile([string]$funcName, [string]$fileName, [string]$title, [string]$description, [string[]]$moduleFuncNames) {
  $out = @(
    'package seed',
    '',
    'import "github.com/pathway/backend/models"',
    '',
    ('func ' + $funcName + '() models.Course {'),
    "`treturn models.Course{",
    ('`t`tTitle:       "' + $title + '",'),
    ('`t`tDescription: "' + $description + '",'),
    "`t`tModules: []models.Module{"
  )

  foreach ($mn in $moduleFuncNames) {
    $out += ("`t`t`t" + $mn + "(),")
  }

  $out += @(
    "`t`t},",
    "`t}",
    '}', ''
  )

  Set-Content -LiteralPath (Join-Path $seedDir $fileName) -Value $out -Encoding utf8
}

function Replace-CourseBlockWithCall([string]$courseTitle, [string]$callName, [int]$endExclusiveIdx) {
  $titleNeedle = 'Title:       "' + $courseTitle + '"'
  $titleIdx = Find-LineIndex $titleNeedle
  if ($null -eq $titleIdx) { throw "Could not find course title: $courseTitle" }

  $start = $titleIdx - 1 # course literal `{` line in slice
  if ($start -lt 0) { throw "Invalid course start for: $courseTitle" }

  $replacement = "`t`t" + $callName + "(),"

  $before = @()
  if ($start -gt 0) { $before = $lines[0..($start - 1)] }

  $after = @()
  if ($endExclusiveIdx -lt $lines.Count) { $after = $lines[$endExclusiveIdx..($lines.Count - 1)] }

  $script:lines = @($before + @($replacement) + $after)
}

Write-Host 'Splitting Design Patterns modules...'
for ($i = 1; $i -le 7; $i++) {
  Write-ModuleFile -moduleId ("patterns-" + $i) -funcName ("modulePatterns" + $i) -fileName ("module_patterns_" + $i + ".go")
}

Write-Host 'Splitting Testing modules...'
for ($i = 1; $i -le 7; $i++) {
  Write-ModuleFile -moduleId ("testing-" + $i) -funcName ("moduleTesting" + $i) -fileName ("module_testing_" + $i + ".go")
}

Write-Host 'Splitting Development Tools modules...'
Write-ModuleFile -moduleId 'vscode-setup' -funcName 'moduleDevTools1' -fileName 'module_dev_tools_1.go'

Write-Host 'Splitting Code Concepts modules...'
$codeModules = @(
  @{ id = 'code-1'; func = 'moduleCodeConcepts1'; file = 'module_code_concepts_1.go' },
  @{ id = 'code-2'; func = 'moduleCodeConcepts2'; file = 'module_code_concepts_2.go' },
  @{ id = 'code-3'; func = 'moduleCodeConcepts3'; file = 'module_code_concepts_3.go' },
  @{ id = 'code-4'; func = 'moduleCodeConcepts4'; file = 'module_code_concepts_4.go' },
  @{ id = 'code-5'; func = 'moduleCodeConcepts5'; file = 'module_code_concepts_5.go' },
  @{ id = 'code-6'; func = 'moduleCodeConcepts6'; file = 'module_code_concepts_6.go' },
  @{ id = 'typescript-starter'; func = 'moduleTypeScriptStarter'; file = 'module_typescript_starter.go' },
  @{ id = 'react-hello-world'; func = 'moduleReactHelloWorld'; file = 'module_react_hello_world.go' },
  @{ id = 'fizzbuzz-class'; func = 'moduleFizzBuzzClass'; file = 'module_fizzbuzz_class.go' },
  @{ id = 'fizzbuzz-galton-ui'; func = 'moduleFizzBuzzGaltonUI'; file = 'module_fizzbuzz_galton_ui.go' }
)
foreach ($m in $codeModules) {
  Write-ModuleFile -moduleId $m.id -funcName $m.func -fileName $m.file
}

Write-Host 'Writing course wrapper files...'
$patternsDesc = Get-CourseDescription 'Design Patterns'
$testingDesc = Get-CourseDescription 'Testing'
$devtoolsDesc = Get-CourseDescription 'Development Tools'
$codeDesc = Get-CourseDescription 'Code Concepts'

Write-CourseFile -funcName 'courseDesignPatterns' -fileName 'course_patterns.go' -title 'Design Patterns' -description $patternsDesc -moduleFuncNames @('modulePatterns1','modulePatterns2','modulePatterns3','modulePatterns4','modulePatterns5','modulePatterns6','modulePatterns7')
Write-CourseFile -funcName 'courseTesting' -fileName 'course_testing.go' -title 'Testing' -description $testingDesc -moduleFuncNames @('moduleTesting1','moduleTesting2','moduleTesting3','moduleTesting4','moduleTesting5','moduleTesting6','moduleTesting7')
Write-CourseFile -funcName 'courseDevelopmentTools' -fileName 'course_dev_tools.go' -title 'Development Tools' -description $devtoolsDesc -moduleFuncNames @('moduleDevTools1')
Write-CourseFile -funcName 'courseCodeConcepts' -fileName 'course_code_concepts.go' -title 'Code Concepts' -description $codeDesc -moduleFuncNames @('moduleCodeConcepts1','moduleCodeConcepts2','moduleCodeConcepts3','moduleCodeConcepts4','moduleCodeConcepts5','moduleCodeConcepts6','moduleTypeScriptStarter','moduleReactHelloWorld','moduleFizzBuzzClass','moduleFizzBuzzGaltonUI')

Write-Host 'Rewriting seed.go to only call course functions...'
$insertIdx = Find-LineIndex '// Insert all courses'
if ($null -eq $insertIdx) { throw 'Could not locate // Insert all courses marker' }

$closeSliceIdx = $null
for ($i = $insertIdx; $i -ge 0; $i--) {
  if ($lines[$i] -match '^\t\}\s*$') { $closeSliceIdx = $i; break }
}
if ($null -eq $closeSliceIdx) { throw 'Could not locate courses slice closing brace' }

# Replace from bottom-up so indices remain stable
Replace-CourseBlockWithCall -courseTitle 'Code Concepts' -callName 'courseCodeConcepts' -endExclusiveIdx $closeSliceIdx
$lines = $script:lines

$ccCallIdx = Find-LineIndex 'courseCodeConcepts(),'
if ($null -eq $ccCallIdx) { throw 'Could not locate courseCodeConcepts() call after replacement' }
Replace-CourseBlockWithCall -courseTitle 'Development Tools' -callName 'courseDevelopmentTools' -endExclusiveIdx $ccCallIdx
$lines = $script:lines

$dtCallIdx = Find-LineIndex 'courseDevelopmentTools(),'
if ($null -eq $dtCallIdx) { throw 'Could not locate courseDevelopmentTools() call after replacement' }
Replace-CourseBlockWithCall -courseTitle 'Testing' -callName 'courseTesting' -endExclusiveIdx $dtCallIdx
$lines = $script:lines

$tCallIdx = Find-LineIndex 'courseTesting(),'
if ($null -eq $tCallIdx) { throw 'Could not locate courseTesting() call after replacement' }
Replace-CourseBlockWithCall -courseTitle 'Design Patterns' -callName 'courseDesignPatterns' -endExclusiveIdx $tCallIdx
$lines = $script:lines

Set-Content -LiteralPath $seedPath -Value $lines -Encoding utf8
Write-Host 'Done.'

